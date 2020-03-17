package manager

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"sync"
	"time"

	"github.com/go-ocf/kit/log"
	"golang.org/x/oauth2"
)

// Manager holds certificates from filesystem watched for changes
type Manager struct {
	mutex          sync.Mutex
	config         clientcredentials.Config
	requestTimeout time.Duration
	token          *oauth2.Token
	tokenErr       error
	doneWg         sync.WaitGroup
	done           chan struct{}
}

// NewManagerFromConfiguration creates a new oauth manager which refreshing token.
func NewManagerFromConfiguration(config Config) (*Manager, error) {
	cfg := config.ToClientCrendtials()
	token, err := getToken(cfg, config.RequestTimeout)
	if err != nil {
		return nil, err
	}
	mgr := &Manager{
		config: cfg,
		token:  token,

		requestTimeout: config.RequestTimeout,
		done:           make(chan struct{}),
	}
	mgr.doneWg.Add(1)

	go mgr.watchToken()

	return mgr, nil
}

// GetToken returns token for clients
func (a *Manager) GetToken(ctx context.Context) (*oauth2.Token, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.token, a.tokenErr
}

// Close ends watching token
func (a *Manager) Close() {
	if a.done != nil {
		close(a.done)
		a.doneWg.Wait()
	}
}

func (a *Manager) nextRenewal() time.Duration {
	t, _ := a.GetToken(context.Background())
	now := time.Now()
	lifetime := t.Expiry.Sub(now) * 2 / 3
	if lifetime < a.requestTimeout {
		lifetime = a.requestTimeout
	}
	return lifetime
}

func getToken(cfg clientcredentials.Config, requestTimeout time.Duration) (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	return cfg.Token(ctx)
}

func (a *Manager) refreshToken() {
	token, err := getToken(a.config, a.requestTimeout)
	if err != nil {
		log.Errorf("cannot refresh token: %v", err)
	}
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.token = token
	a.tokenErr = err
}

func (a *Manager) watchToken() {
	defer a.doneWg.Done()
	for {
		nextRenewal := a.nextRenewal()
		select {
		case <-a.done:
			return
		case <-time.After(nextRenewal):
			a.refreshToken()
		}
	}
}
