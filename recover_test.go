package recover

import (
	"testing"

	"errors"
)

func TestString(t *testing.T) {
	var err error
	defer func() {
		if err.Error() != "test panic" {
			t.Errorf("Recover() = %+v, expected = %+v", err, "test panic")
		}
	}()
	defer Recover(&err)
	panic("test panic")
}

func TestError(t *testing.T) {
	err := errors.New("test error")
	var perr error
	defer func() {
		if perr != err {
			t.Errorf("Recover() = %v, expected = %v", err, perr)
		}
	}()
	defer Recover(&perr)
	panic(err)
}

func TestDefault(t *testing.T) {
	var err error
	defer func() {
		if err.Error() != "unexpected Panic: 123" {
			t.Errorf("Recover() = %v, expected = %v", err, "unexpected Panic: 123")
		}
	}()
	defer Recover(&err)
	panic(123)
}
