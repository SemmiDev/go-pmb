package domain

import "github.com/myesui/uuid"

type ID string

func (i ID) String() string {
	return string(i)
}

func GenerateID() ID {
	return ID(uuid.NewV4().String())
}

func (i ID) Empty() bool {
	return string(i) == ""
}
