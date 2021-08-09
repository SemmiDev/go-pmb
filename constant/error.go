package constant

import "errors"

var (
	ErrEmail              = errors.New("e-mail has been registered")
	ErrPhone              = errors.New("phone has been registered")
	ErrIDNotFound         = errors.New("ID not found")
	ErrRegisterNotFound   = errors.New("Register not found")
	ErrBillHasNotBeenPaid = errors.New("Please pay the bill first")
)
