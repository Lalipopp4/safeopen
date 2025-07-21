package safeopen

import (
	"io"
	"runtime"
)

func Open[T io.Closer](cl T, err error) (T, error) {
	type data struct {
		res T
	}

	runtime.AddCleanup(&data{cl}, func(cl T) {
		cl.Close()
	}, cl)

	return cl, err
}
