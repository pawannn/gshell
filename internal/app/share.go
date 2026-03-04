package app

import (
	"fmt"

	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/transport"
	"github.com/pawannn/gshell/pkg"
)

func Share(port string, sessionName string) error {

	ip, err := pkg.GetLocalIP()
	if err != nil {
		return err
	}

	payload := security.SessionPayload{
		IP:      ip,
		Port:    port,
		Session: sessionName,
	}

	token, err := security.Encrypt(payload)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Session started:", sessionName)
	fmt.Println("Invite command:")
	fmt.Println("gshell join", token)
	fmt.Println()

	return transport.StartListener(port)
}
