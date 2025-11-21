package must

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Must panicked unexpectedly: %v", r)
		}
	}()
	Must(nil)
}

func TestMustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Must did not panic")
		}
	}()
	Must(errors.New("error"))
}

func TestMustV(t *testing.T) {
	val := 10
	res := MustV(val, nil)
	if res != val {
		t.Errorf("expected %d, got %d", val, res)
	}
}

func TestMustVPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustV did not panic")
		}
	}()
	MustV(0, errors.New("error"))
}
