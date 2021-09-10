package token

import "time"

type Maker interface {
	CreateToken(u string, d time.Duration) (string, error)
	VerifyToken(t string) (*Payload, error)
}
