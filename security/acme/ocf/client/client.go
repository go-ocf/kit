package client

import (
	"fmt"

	"github.com/go-acme/lego/challenge/resolver"
	"github.com/go-acme/lego/lego"
	"github.com/go-acme/lego/registration"
)

type Client struct {
	c        *lego.Client
	deviceID string
}

type Config struct {
	lego.Config
	DeviceID string
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.DeviceID == "" {
		return nil, fmt.Errorf("invalid DeviceID")
	}

	c, err := lego.NewClient(&cfg.Config)
	if err != nil {
		return nil, err
	}
	return &Client{
		c: c,
	}, nil
}

func (c *Client) Certificate() *certifier {
	return &certifier{c: c.c.Certificate, deviceID: c.deviceID}
}

func (c *Client) Challenge() *resolver.SolverManager {
	return c.c.Challenge
}

func (c *Client) Registration() *registration.Registrar {
	return c.c.Registration
}
