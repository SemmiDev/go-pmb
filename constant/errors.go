package constant

import "errors"

var (
	ErrEmail              = errors.New("e-mail has been registered")
	ErrPhone              = errors.New("phone has been registered")
	ErrIDNotFound         = errors.New("ID not found")
	ErrRegisterNotFound   = errors.New("register not found")
	ErrBillHasNotBeenPaid = errors.New("please pay the bill first")
)
