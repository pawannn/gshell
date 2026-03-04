package transport

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func StartListener(port string) error {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	defer ln.Close()
	log.Println("Listening on port", port)

	conn, err := ln.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()
	fmt.Println("Client connected from : ", conn.RemoteAddr())

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)

	return nil
}
