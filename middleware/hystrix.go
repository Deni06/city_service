package middleware

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"

	"context"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               10,
	})
	return hystrix.Do(req.Service()+"."+req.Method(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, nil)
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
