package gerr

import (
	"fmt"
	"io"
)

type CodeErr struct {
	code  int    // 错误码
	msg   string // 错误信息前端友好提升
	cause error  // 报错源头错误，不应该给前端看的信息
}

func (e *CodeErr) Error() string {
	return e.msg
}
func (e *CodeErr) Code() int {
	return e.code
}
func (e *CodeErr) Cause() error { return e.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *CodeErr) Unwrap() error { return e.cause }
func (e *CodeErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v \n", e.Cause())
			_, _ = io.WriteString(s, e.msg)
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, e.Error())
	}
}
func wrapErrf(err error, code int, format string, a ...any) Error {
	e := &CodeErr{
		code:  code,
		msg:   fmt.Sprintf(format, a...),
		cause: err,
	}
	return &withStack{
		CodeErr: e,
		stack:   callers(4),
	}
}

// baseErr 系统错误
// code: 错误码
// stackSkip 栈起始记录的位置
func baseErr(msg string, code int) Error {
	e := &CodeErr{
		code: code,
		msg:  msg,
	}
	return &withStack{
		CodeErr: e,
		stack:   callers(4),
	}
}

func New(code int, msg string) Coder {
	return func() (int, string) {
		return code, msg
	}
}
