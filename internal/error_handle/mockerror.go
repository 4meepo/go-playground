package main

import (
	"github.com/pkg/errors"
	"os"
)

func MustError() error {
	_, err := os.Open("/user/sdfsdf")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
