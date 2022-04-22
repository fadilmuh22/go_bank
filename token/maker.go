package token

import (
	"fmt"
	"time"

	"golang.org/x/crypto/chacha20poly1305"
)

type Payload struct{}

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

func NewPasetoMaker(symmetrickKey string) (Maker, error) {

	if len(symmetrickKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return nil, nil

}
