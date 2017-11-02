package recover

import (
	"github.com/pkg/errors"
)

// Recover 获取Panicking信息，转换为error
func Recover(err *error) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			*err = errors.New(x)
		case error:
			*err = x
		default:
			*err = errors.Errorf("unexpected Panic: %v", x)
		}
	}
}
