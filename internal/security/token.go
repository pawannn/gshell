package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
)

func Encrypt(payload SessionPayload) (string, error) {

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	key := []byte(SecretKey)

	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(token string) (*SessionPayload, error) {

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	key := []byte(SecretKey)

	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	var payload SessionPayload
	err = json.Unmarshal(plaintext, &payload)

	return &payload, err
}
