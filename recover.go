package recover

import "github.com/pkg/errors"

// Recover 获取Panicking信息，转换为error
func Recover(err *error) {
	if r := recover(); r != nil {
		var ok bool
		*err, ok = r.(error)
		if !ok {
			*err = errors.Errorf("unexpected Panic: %v", r)
		}
		*err = errors.Wrap(*err, "unexpected error")
	}
}
