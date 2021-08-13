package e

import (
	"fmt"
	"github.com/pkg/errors"
)

type errno struct {
	code int
	msg  string
}

func (e errno) Error() string {
	return e.msg
}

type codeError struct {
	code int
	msg  string
	err  error
}

func (e *codeError) Error() string {
	return fmt.Sprintf("error: code = %d msg = %s err = %s", e.code, e.msg, e.err)
}

// New 新建错误
func New(errno *errno, err error) *codeError {
	return &codeError{code: errno.code, msg: errno.msg, err: err}
}

// Add 添加错误消息
func (e *codeError) Add(msg string) *codeError {
	e.msg += " " + msg
	return e
}

// Addf 添加错误消息
func (e *codeError) Addf(format string, args ...interface{}) *codeError {
	e.msg += " " + fmt.Sprintf(format, args...)
	return e
}

// NewWithStack 新建错误，附加调用栈
func NewWithStack(errno *errno, err error) error {
	return errors.WithStack(&codeError{code: errno.code, msg: errno.msg, err: err})
}

// Wrap 包装错误信息，附加调用栈
// 第二个参数只能是 string，也可以不传，大部分情况不用传
func Wrap(err error, args ...interface{}) error {
	if len(args) >= 1 {
		if msg, ok := args[0].(string); ok {
			return errors.Wrap(err, msg)
		}
	}

	return errors.Wrap(err, "")
}

// Cause 获取原始错误对象
func Cause(err error) error {
	return errors.Cause(err)
}

// DecodeErr 提取错误码，返回 code 和 msg
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.code, OK.msg
	}
	switch typed := err.(type) {
	case *errno:
		return typed.code, typed.msg
	case *codeError:
		return typed.code, typed.msg
	}
	return InternalServerError.code, err.Error()
}
