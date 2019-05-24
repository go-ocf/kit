package security

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
)

// Generates `func IsInsecure() bool`
//go:generate go run generateInsecure.go security

// TLSConfig set configuration.
type TLSConfig struct {
	Certificate    string `envconfig:"TLS_CERTIFICATE"`
	CertificateKey string `envconfig:"TLS_CERTIFICATE_KEY"`
	CAPool         string `envconfig:"TLS_CA_POOL"`
}

// SetTLSConfig setup tls.Config that provides verification certificate with connection.
func SetTLSConfig(config TLSConfig, certificateVerifier CertificateVerifier) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(config.Certificate, config.CertificateKey)
	if err != nil {
		return nil, fmt.Errorf("cannot load x509 key pair('%v', '%v'): %v", config.Certificate, config.CertificateKey, err)
	}

	caRootPool := x509.NewCertPool()
	caIntermediatesPool := x509.NewCertPool()

	err = filepath.Walk(config.CAPool, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return fmt.Errorf("cannot walk through directory '%v': %v", config.CAPool, e)
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			certPEMBlock, err := ioutil.ReadFile(path)
			if err != nil {
				return nil
			}
			certDERBlock, _ := pem.Decode(certPEMBlock)
			if certDERBlock == nil {
				return nil
			}
			if certDERBlock.Type != "CERTIFICATE" {
				return nil
			}
			caCert, err := x509.ParseCertificate(certDERBlock.Bytes)
			if err != nil {
				return nil
			}
			if bytes.Compare(caCert.RawIssuer, caCert.RawSubject) == 0 && caCert.IsCA {
				caRootPool.AddCert(caCert)
			} else if caCert.IsCA {
				caIntermediatesPool.AddCert(caCert)
			} else {
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(caRootPool.Subjects()) == 0 {
		return nil, fmt.Errorf("CA Root pool is empty")
	}

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    caRootPool,
		RootCAs:      caRootPool,

		GetConfigForClient: func(info *tls.ClientHelloInfo) (*tls.Config, error) {
			//https://github.com/golang/go/issues/29895
			m := tls.Config{
				Certificates: []tls.Certificate{cert},
				ClientAuth:   tls.RequireAnyClientCert,
			}
			m.VerifyPeerCertificate = newVerifyPeerCert(caIntermediatesPool, caRootPool, info.Conn, certificateVerifier)
			return &m, nil
		},
	}

	return &tlsConfig, nil
}

func newVerifyPeerCert(intermediates *x509.CertPool, roots *x509.CertPool, conn net.Conn, certificateVerifier CertificateVerifier) func(rawCerts [][]byte, verifyChains [][]*x509.Certificate) error {
	return func(rawCerts [][]byte, verifyChains [][]*x509.Certificate) error {
		for _, rawCert := range rawCerts {
			cert, err := x509.ParseCertificate(rawCert)
			if err != nil {
				return err
			}

			_, err = cert.Verify(x509.VerifyOptions{
				Intermediates: intermediates,
				Roots:         roots,
				CurrentTime:   time.Now(),
				KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
			})
			if err != nil {
				return err
			}
			err = certificateVerifier.Verify(conn, cert)
			if err != nil {
				return err
			}
		}
		return nil
	}
}