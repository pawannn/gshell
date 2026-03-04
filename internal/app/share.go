package app

import (
	"fmt"
	"time"

	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/transport"
	"github.com/pawannn/gshell/pkg"
)

func Share(port string, sessionName string, password string) error {

	ip, err := pkg.GetLocalIP()
	if err != nil {
		return err
	}

	payload := security.SessionPayload{
		IP:      ip,
		Port:    port,
		Session: sessionName,
		Pwd:     password,
		Expiry:  time.Now().Add(30 * time.Second).Unix(),
	}

	token, err := security.Encrypt(payload)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Session:", sessionName)
	fmt.Println("Token Expires in : 30 seconds")
	fmt.Println("Join using:")
	fmt.Println("gshell join", token)
	fmt.Println()

	return transport.StartListener(payload)
}
