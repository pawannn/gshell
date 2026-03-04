package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/transport"
)

func Join(token string, password string) error {
	payload, err := security.Decrypt(token)
	if err != nil {
		return err
	}

	if time.Now().Unix() > payload.Expiry {
		return errors.New("session invite expired")
	}

	if payload.Pwd != password {
		return errors.New("invalid password")
	}

	fmt.Println("Joining session:", payload.Session)

	return transport.Connect(*payload)
}
