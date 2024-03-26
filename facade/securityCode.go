package facade

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (c *SecurityCode) checkSecurityCode(code int) error {
	if c.code != code {
		return errors.New(fmt.Sprintf(" security code %d is not a valid security code", code))
	}
	return nil
}
