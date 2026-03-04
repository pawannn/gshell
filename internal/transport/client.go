package transport

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pawannn/gshell/internal/security"
	"golang.org/x/term"
)

func Connect(payload security.SessionPayload) error {

	address := payload.IP + ":" + payload.Port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Connected to", payload.Session)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	stdio := struct {
		io.Reader
		io.Writer
	}{
		os.Stdin,
		os.Stdout,
	}

	Pipe(conn, stdio)

	return nil
}
