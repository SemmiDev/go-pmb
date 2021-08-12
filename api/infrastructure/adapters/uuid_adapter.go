package adapters

import "github.com/twinj/uuid"

type IUuidAdapter interface {
	GenerateUuid() string
}

type UuidAdapter struct {
}

func NewUuidAdapter() IUuidAdapter {
	return &UuidAdapter{}
}

func (u *UuidAdapter) GenerateUuid() string {
	return uuid.NewV4().String()
}
