package data

import "github.com/pkg/errors"

var (
	ErrRowExist = errors.New("The row in table exist.")
)
