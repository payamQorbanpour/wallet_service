package model

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	CodeErrUnexpected         = 0
	CodeErrWalletExists       = 1
	CodeErrWalletDoesNotExist = 2
	CodeErrNotEnoghBalance    = 3
	CodeErrSelfTransaction    = 4
)

var (
	ErrUnexpected         = errors.New("unexpected error")
	ErrWalletExists       = errors.New("wallet with this phone number already exists")
	ErrWalletDoesNotExist = errors.New("wallet with this phone number does not exist")
	ErrNotEnoghBalance    = errors.New("there is not enough balance to commit the transaction")
	ErrSelfTransaction    = errors.New("transaction to self wallet is not possible")
)

func ErrToCode(err error) int {
	switch errors.Cause(err) {
	case ErrWalletExists:
		return CodeErrWalletExists
	case ErrWalletDoesNotExist:
		return CodeErrWalletDoesNotExist
	case ErrNotEnoghBalance:
		return CodeErrNotEnoghBalance
	case ErrSelfTransaction:
		return CodeErrSelfTransaction
	default:
		return CodeErrUnexpected
	}
}

func ErrToHTTPStatus(err error) int {
	switch errors.Cause(err) {
	case ErrWalletExists, ErrWalletDoesNotExist, ErrNotEnoghBalance, ErrSelfTransaction:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
