package model

import (
	"github.com/SemmiDev/fiber-go-clean-arch/constant"
)

type (
	RegistrationRequest struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Program string `json:"program"`
	}

	RegistrationResponse struct {
		Recipient      string        `json:"-"`
		Username       string        `json:"username"`
		Password       string        `json:"password"`
		Bill           constant.Bill `json:"bill"`
		VirtualAccount string        `json:"virtual_account"`
	}

	UpdateStatus struct {
		VirtualAccount string `json:"virtual_account"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
