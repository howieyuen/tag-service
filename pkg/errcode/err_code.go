package errcode

import (
	"fmt"
)

type Error struct {
	code int
	msg  string
}

var _code = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, flag := _code[code]; flag {
		panic(fmt.Sprintf("code %d already exist，please change to annother", code))
	}
	_code[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.Code(), e.Msg())
}

func (e Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
