package requests

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Login) IsValid() error {
	if r.Username == "" {
		return customErrors.UsernameIsRequired
	}

	if r.Password == "" {
		return customErrors.PasswordIsRequired
	}

	return nil
}
