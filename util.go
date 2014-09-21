package jnigo

import (
	"errors"
	"fmt"
)

func Err(f string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(f, args...))
}
