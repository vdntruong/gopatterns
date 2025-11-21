package pointer

import (
	"testing"
)

func TestPointerOf(t *testing.T) {
	val := 10
	ptr := PointerOf(val)

	if ptr == nil {
		t.Fatal("expected pointer to be not nil")
	}
	if *ptr != val {
		t.Errorf("expected %d, got %d", val, *ptr)
	}

	strVal := "hello"
	strPtr := PointerOf(strVal)
	if strPtr == nil {
		t.Fatal("expected pointer to be not nil")
	}
	if *strPtr != strVal {
		t.Errorf("expected %s, got %s", strVal, *strPtr)
	}
}
