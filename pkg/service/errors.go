package service

import "errors"

var (
	ErrUserDoesNotExist              = errors.New("user does not exist")
	ErrNotEnoughMoney                = errors.New("not enough money")
	ErrIncorrectParameters           = errors.New("incorrect parameters")
	ErrReservedPurchaseAlreadyExists = errors.New("reserved purchase with these settings is already exists")
	ErrReservedPurchaseDoesNotExist  = errors.New("reserved purchase with these settings does not exist")
)
