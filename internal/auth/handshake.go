package auth

import (
	"encoding/json"
	"errors"
	"net"
)

type Handshake struct {
	Password string `json:"password"`
}

func ServerHandshake(conn net.Conn, expected string) error {

	var msg Handshake
	err := json.NewDecoder(conn).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Password != expected {
		return errors.New("authentication failed")
	}

	return nil
}

func ClientHandshake(conn net.Conn, password string) error {

	msg := Handshake{
		Password: password,
	}

	return json.NewEncoder(conn).Encode(msg)
}
