package a

import "errors"

func f() {
	err := errors.New("error")
	_ = err == nil                // OK
	_ = err == err                // want "must use errors.Is"
	_, _ = err.(interface{ F() }) // want "must use errors.As"
}
