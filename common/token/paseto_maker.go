package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidKeySize = "invalid key size: must be exactly %d characters"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (maker *PasetoMaker) CreateToken(u string, d time.Duration) (string, error) {
	p, err := NewPayload(u, d)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetricKey, p, nil)
}

func (maker *PasetoMaker) VerifyToken(t string) (*Payload, error) {
	p := new(Payload)
	err := maker.paseto.Decrypt(t, maker.symmetricKey, p, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = p.Valid()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewPasetoMaker(sk string) (Maker, error) {
	if len(sk) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(ErrInvalidKeySize, chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(sk),
	}
	return maker, nil
}
