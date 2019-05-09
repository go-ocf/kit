package nats

import (
	cqrsEventBus "github.com/go-ocf/cqrs/eventbus"
	cqrsNats "github.com/go-ocf/cqrs/eventbus/nats"
	cqrsUtils "github.com/go-ocf/kit/cqrs"
)

type Subscriber struct {
	*cqrsNats.Subscriber
}

// NewSubscriber create new subscriber with proto unmarshaller.
func NewSubscriber(config Config, goroutinePoolGo cqrsEventBus.GoroutinePoolGoFunc, errFunc cqrsEventBus.ErrFunc) (*Subscriber, error) {
	s, err := cqrsNats.NewSubscriber(config.URL, cqrsUtils.Unmarshal, goroutinePoolGo, errFunc, config.Options...)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		s,
	}, nil
}

// Close closes the publisher.
func (p *Subscriber) Close() {
	p.Subscriber.Close()
}
