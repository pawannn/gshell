package transport

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	defer conn.Close()

	fmt.Println("Connected to", address)

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
	return nil
}
