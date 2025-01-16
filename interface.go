package gerr

import "fmt"

type Error interface {
	Error() string

	Code() int

	Cause() error

	Unwrap() error

	Format(s fmt.State, verb rune)
}

//type ErrorLog interface {
//	Error(err error)
//}

type Coder func() (code int, info string)
