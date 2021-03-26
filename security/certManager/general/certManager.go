package general

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/plgd-dev/kit/security"
	"go.uber.org/zap"
)

// Config provides configuration of a file based Server Certificate manager
type Config struct {
	CAFile                    string `yaml:"caFile" json:"caFile" envconfig:"CA_FILE" description:"file path to the root certificate in PEM format"`
	KeyFile                   string `yaml:"keyFile" json:"keyFile" envconfig:"KEY_FILE" description:"file name of private key in PEM format"`
	CertFile                  string `yaml:"certFile" json:"certFile" envconfig:"CERT_FILE" description:"file name of certificate in PEM format"`
	ClientCertificateRequired bool   `yaml:"clientCertificateRequired" json:"clientCertificateRequired" envconfig:"CLIENT_CERTIFICATE_REQUIRED" description:"require client certificate"`
	UseSystemCAPool           bool   `yaml:"useSystemCAPool" json:"useSystemCAPool" envconfig:"USE_SYSTEM_POOL" description:"use system certification pool"`
}

func (c Config) Validate() error {
	if c.CAFile == "" && !c.UseSystemCAPool {
		return fmt.Errorf("invalid caFile")
	}
	if c.CertFile == "" {
		return fmt.Errorf("invalid certFile")
	}
	if c.KeyFile == "" {
		return fmt.Errorf("invalid keyFile")
	}
	return nil
}

func (c *Config) SetDefaults() {
	c.ClientCertificateRequired = true
}

// CertManager holds certificates from filesystem watched for changes
type CertManager struct {
	config Config

	watcher                 *fsnotify.Watcher
	doneWg                  sync.WaitGroup
	done                    chan struct{}
	verifyClientCertificate tls.ClientAuthType
	logger                  *zap.Logger

	mutex      sync.Mutex
	tlsKeyPair tls.Certificate
	tlsCAPool  *x509.CertPool
}

// New creates a new certificate manager which watches for certs in a filesystem
func New(config Config, logger *zap.Logger) (*CertManager, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	verifyClientCertificate := tls.RequireAndVerifyClientCert
	if !config.ClientCertificateRequired {
		verifyClientCertificate = tls.NoClientCert
	}

	c := &CertManager{
		watcher:                 watcher,
		config:                  config,
		verifyClientCertificate: verifyClientCertificate,
		logger:                  logger,
	}
	err = c.loadCAs()
	if err != nil {
		return nil, err
	}
	err = c.loadCerts()
	if err != nil {
		return nil, err
	}
	if err := c.watcher.Add(filepath.Dir(config.CAFile)); err != nil {
		return nil, err
	}
	if err := c.watcher.Add(filepath.Dir(config.CertFile)); err != nil {
		return nil, err
	}
	if err := c.watcher.Add(filepath.Dir(config.KeyFile)); err != nil {
		return nil, err
	}

	c.done = make(chan struct{})
	c.doneWg.Add(1)

	go c.watchFiles()

	return c, nil
}

// GetCertificateAuthorities returns certificates authorities
func (a *CertManager) GetCertificateAuthorities() *x509.CertPool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.tlsCAPool
}

// GetServerTLSConfig returns tls configuration for servers
func (a *CertManager) GetServerTLSConfig() *tls.Config {
	return &tls.Config{
		ClientCAs:      a.GetCertificateAuthorities(),
		GetCertificate: a.getServerCertificate,
		MinVersion:     tls.VersionTLS12,
		ClientAuth:     a.verifyClientCertificate,
	}
}

// GetClientTLSConfig returns tls configuration for clients
func (a *CertManager) GetClientTLSConfig() *tls.Config {
	return &tls.Config{
		RootCAs:                  a.GetCertificateAuthorities(),
		GetClientCertificate:     a.getClientCertificate,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}
}

// Close ends watching certificates
func (a *CertManager) Close() {
	if a.done != nil {
		_ = a.watcher.Close()
		close(a.done)
		a.doneWg.Wait()
	}
}

func (a *CertManager) getServerCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return &a.tlsKeyPair, nil
}

func (a *CertManager) getClientCertificate(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return &a.tlsKeyPair, nil
}

func (a *CertManager) loadCerts() error {
	if a.config.KeyFile != "" && a.config.CertFile != "" {
		keyPath := a.config.KeyFile
		tlsKey, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return fmt.Errorf("cannot load certificate key from '%v': %w", keyPath, err)
		}
		certPath := a.config.CertFile
		tlsCert, err := ioutil.ReadFile(certPath)
		if err != nil {
			return fmt.Errorf("cannot load certificate from '%v': %w", certPath, err)
		}
		cert, err := tls.X509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return fmt.Errorf("cannot load certificate pair: %w", err)
		}
		a.setTLSKeyPair(cert)
		return nil
	}
	return fmt.Errorf("cannot load certificate from CertFile='%v', KeyFile='%v'", a.config.CertFile, a.config.KeyFile)

}
func (a *CertManager) loadCAs() error {
	var cas []*x509.Certificate
	if a.config.CAFile != "" {
		certs, err := security.LoadX509(a.config.CAFile)
		if err != nil {
			return fmt.Errorf("cannot load certificate authorities from '%v': %w", a.config.CAFile, err)
		}
		cas = certs
	}
	if a.config.UseSystemCAPool {
		a.setCAPool(security.NewDefaultCertPool(cas))
	} else {
		p := x509.NewCertPool()
		for _, c := range cas {
			p.AddCert(c)
		}
		a.setCAPool(p)
	}
	return nil
}

func (a *CertManager) setTLSKeyPair(cert tls.Certificate) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.tlsKeyPair = cert
}

func (a *CertManager) setCAPool(capool *x509.CertPool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.tlsCAPool = capool
}

func (a *CertManager) watchFiles() {
	defer a.doneWg.Done()
	for {
		var updateCert, updateKey, updateCAs bool
		select {
		case <-a.done:
			return
			// watch for events
		case event := <-a.watcher.Events:
			switch event.Op {
			case fsnotify.Create, fsnotify.Write, fsnotify.Rename:
				if strings.Contains(event.Name, a.config.KeyFile) {
					updateKey = true
				}

				if strings.Contains(event.Name, a.config.CertFile) {
					updateCert = true
				}

				if strings.Contains(event.Name, a.config.CAFile) {
					updateCAs = true
				}
			}
		}
		if updateCert && updateKey {
			err := a.loadCerts()
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
		if updateCAs {
			err := a.loadCAs()
			if err != nil {
				a.logger.Error(err.Error())
			}
		}
	}
}
