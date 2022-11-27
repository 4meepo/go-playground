package learn_golang

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestFunctionalOption(t *testing.T) {
	opts := []Option{WithAddress("local"), WithPort(80)}

	// var o options
	var o options
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			panic(err)
		}
	}
	fmt.Println(*o.address, *o.port)

}

type options struct {
	address *string
	port    *uint
}

type Option func(o *options) error

func WithAddress(a string) Option {
	return func(o *options) error {
		if a == "" {
			return errors.New("....")
		}
		o.address = &a
		return nil
	}
}

func WithPort(p uint) Option {
	return func(o *options) error {
		o.port = &p
		return nil
	}
}
