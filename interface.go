package gerr

import (
	"context"
	"fmt"
)

type Error interface {
	Error() string

	Code() int

	Cause() error

	Unwrap() error

	Format(s fmt.State, verb rune)

	Log(logger Logger)

	LogCtx(ctx context.Context, logger Logger)
}

//type ErrorLog interface {
//	Error(err error)
//}

type Coder func() (code int, info string)

type Logger interface {
	Error(err error)
	ErrorCtx(ctx context.Context, err error)
}
