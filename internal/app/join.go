package app

import (
	"fmt"

	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/transport"
)

func Join(token string) error {

	payload, err := security.Decrypt(token)
	if err != nil {
		return err
	}

	fmt.Println("Joining session:", payload.Session)

	return transport.Connect(*payload)
}
