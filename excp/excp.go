package excp

import (
	"fmt"
	"runtime/debug"
)

func Try(f func(), exception *error) {
	if exception == nil {
		panic("exception is nil")
	}
	defer func() {
		if err := recover(); err != nil {
			switch e := err.(type) {
			case error:
				*exception = e
			default:
				*exception = NewDefaultError(err)
			}
		}
	}()
	f()
}

func TryWithStack(f func(), exception *error,stack *[]byte) {
	if exception == nil {
		panic("exception is nil")
	}
	defer func() {
		if err := recover(); err != nil {
			*stack = debug.Stack()
			switch e := err.(type) {
			case error:
				*exception = e
			default:
				*exception = NewDefaultError(err)
			}
		}
	}()
	f()
}




func TryCatch(try func(), catch func(err error)) {
	defer func() {
		if err := recover(); err != nil {
			switch e := err.(type) {
			case error:
				catch(e)
			default:
				catch(NewDefaultError(err))
			}
		}
	}()
	try()
}

func TryCatchWithStack(try func(), catch func(err error, stack []byte)) {
	defer func() {
		if err := recover(); err != nil {
			switch e := err.(type) {
			case error:
				catch(e,debug.Stack())
			default:
				catch(NewDefaultError(err),debug.Stack())
			}
		}
	}()
	try()
}



func Throw(e error) {
	panic(e)
}

type DefaultError struct {
	E interface{}
}

func (d DefaultError) Error() string {
	return fmt.Sprintf("%v", d.E)
}

func NewDefaultError(v interface{}) error {
	return &DefaultError{
		E: v,
	}
}

type WithStack struct {
	Stack []byte
	E     error
}

func (w WithStack) Error() string {
	return w.E.Error()
}

func WithStackError(e error) {

}