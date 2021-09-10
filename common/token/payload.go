package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(u string, d time.Duration) (*Payload, error) {
	tID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	p := &Payload{
		ID:        tID,
		Username:  u,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(d),
	}
	return p, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
