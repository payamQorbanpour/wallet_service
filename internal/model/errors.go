package model

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	CodeErrUnexpected         = 0
	CodeErrWalletExists       = 1
	CodeErrWalletDoesNotExist = 2
)

var (
	ErrUnexpected         = errors.New("unexpected error")
	ErrWalletExists       = errors.New("wallet with this phone number already exists")
	ErrWalletDoesNotExist = errors.New("wallet with this phone number does not exist")
)

func ErrToCode(err error) int {
	switch errors.Cause(err) {
	case ErrWalletExists:
		return CodeErrWalletExists
	case ErrWalletDoesNotExist:
		return CodeErrWalletDoesNotExist
	default:
		return CodeErrUnexpected
	}
}

func ErrToHTTPStatus(err error) int {
	switch errors.Cause(err) {
	case ErrWalletExists, ErrWalletDoesNotExist:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
