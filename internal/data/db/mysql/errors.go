package mysql

import "github.com/pkg/errors"

var (
	ErrNotFound   = errors.New("Item not found in table")
	ErrValidation = errors.New("Valid column error")
)
