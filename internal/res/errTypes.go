package res

import (
	"errors"
	"fmt"
)

var badRequest = errors.New("bad Request")

var internal = errors.New("internal error")

func wrapType(err error, typ error, text string) error {
	return fmt.Errorf("%w, %s(%s)", typ, text, err.Error())
}

func BadRequestErr1(err error) error {
	return wrapType(err, badRequest, "")
}

func InternalErr1(err error) error {
	return wrapType(err, internal, "")
}

func BadRequestErr2(err error, text string) error {
	return wrapType(err, badRequest, text)
}

func InternalErr2(err error, text string) error {
	return wrapType(err, internal, text)
}
