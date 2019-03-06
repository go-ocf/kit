package net

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTLSListener_AcceptContext(t *testing.T) {
	ctxCanceled, ctxCancel := context.WithCancel(context.Background())
	ctxCancel()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
			},
		},
		{
			name: "cancelled",
			args: args{
				ctx: ctxCanceled,
			},
			wantErr: true,
		},
	}

	dir, err := ioutil.TempDir("", "gotesttmp")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)
	TLSConfig := testSetupTLS(t, dir)

	config, err := SetTLSConfig(TLSConfig, func(conn net.Conn, certificate *x509.Certificate) error {
		return nil
	})
	assert.NoError(t, err)

	listener, err := NewTLSListener("tcp", "127.0.0.1:", config, time.Millisecond*100)
	assert.NoError(t, err)
	defer listener.Close()

	go func() {
		for i := 0; i < len(tests); i++ {
			cert, err := tls.X509KeyPair(CertPEMBlock, KeyPEMBlock)
			assert.NoError(t, err)

			c, err := tls.Dial("tcp", listener.Addr().String(), &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
			})
			assert.NoError(t, err)
			_, err = c.Write([]byte("hello"))
			assert.NoError(t, err)

			time.Sleep(time.Millisecond * 200)
			c.Close()
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con, err := listener.AcceptContext(tt.args.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				b := make([]byte, 1024)
				_, err = con.Read(b)
				assert.NoError(t, err)
				err = con.Close()
				assert.NoError(t, err)
			}
		})
	}
}

var (
	// CertPEMBlock is a X509 data used to test TLS servers (used with tls.X509KeyPair)
	CertListenerPEMBlock = []byte(`-----BEGIN CERTIFICATE-----
MIICETCCAXqgAwIBAgIQGncx7Aoc6cmxB0O2AlDbIjANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCB
iQKBgQDNHitCs03rbqjQq77c6mlgNX68mew9Mn030JnHLhgWblGaMUsMqUPJn7Lx
i5BPnlc7rIEUHhhV38WmjSgQ7nvkZBM4A6lyyR3B3Vk+rQw6Xukj/ix+BXGoMZM9
sZFj4XZr+9n0ocXNSk3d+b43Ug42q5W17WYm10t2/ZYBkH9ISQIDAQABo2YwZDAO
BgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUw
AwEB/zAsBgNVHREEJTAjgglsb2NhbGhvc3SHBH8AAAGHEAAAAAAAAAAAAAAAAAAA
AAEwDQYJKoZIhvcNAQELBQADgYEAdqjf/9CuyOjgdwMAb1k3lO9+lwWr6dq0zXwU
zq0Qj5spgLxeRK+SRwSswW2VbszkSr+Qd4OVDlX10KCzBZJ5qRZWcwM755UPxd+e
oO0RFbASO4yrMduKkXJo6tiMS/rjEC+9yUEEltlZduuQqIAdDjvgZfmhfMQpNuD/
X6zS+rU=
-----END CERTIFICATE-----`)

	// KeyPEMBlock is a X509 data used to test TLS servers (used with tls.X509KeyPair)
	KeyListenerPEMBlock = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDNHitCs03rbqjQq77c6mlgNX68mew9Mn030JnHLhgWblGaMUsM
qUPJn7Lxi5BPnlc7rIEUHhhV38WmjSgQ7nvkZBM4A6lyyR3B3Vk+rQw6Xukj/ix+
BXGoMZM9sZFj4XZr+9n0ocXNSk3d+b43Ug42q5W17WYm10t2/ZYBkH9ISQIDAQAB
AoGBAJXiEriFr013KjJ5HVnujJu522dTjnXVe/yaGJScUQurB0QF+xJAaYFeifLJ
CeW0DYhUcGnT5/JwNsySXxGoQqx8QCfStH8c6ZPkAF3qXYbPNsX4x2IpDJYyp7ve
Qj501VpeRPNd3mueBHvkZ0UPkBo6Tz7iA6ilp5qgF2soMUsBAkEA0Mwu4NSNRf7u
Gg42U9aFa0y9TZ5QuKLC42+SwzbtTyfMSj5G+m05aeuqinmWhNesaBss4BmmmSXg
J0N6kekUaQJBAPt9Bt1pJPKGv6IbC3SsccooRS9sQOUhOTRiVnwzZ1i4Dk23fRQN
Rox2AzYzsMPG6vGRwumQuBvj6RZy+BGWmOECQBF82HxKMR7osCaMhC5XbEtFXSGQ
HfCo6SvFX4RsKEoV6j1Zo/Y7ibB+ZYU9k8bCjZUWmZaXb2WqT3DabPyliekCQQDN
UUDGiO4KNurDLPNIWPU5h3Eci3Pb3Sj31IUpN0pbi0DaQECUm1YKnNp4aPEalQ8B
E/CegXFeC88jbc+LhHjhAkEAv6N2yaaKphaFOYLdcApVViIwKfdoZFKm+hEikhHg
zlI1KSI23j1bIvJXxH2sWMhbu534p3rE1MqC6v5dc/dGZA==
-----END RSA PRIVATE KEY-----`)
)