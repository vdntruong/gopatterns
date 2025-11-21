package pointer_test

import (
	"fmt"

	"github.com/vdntruong/gopatterns/pkg/pointer"
)

func ExamplePointerOf() {
	// Get a pointer to an integer literal
	intPtr := pointer.PointerOf(42)
	fmt.Printf("Int: %d\n", *intPtr)

	// Get a pointer to a string literal
	strPtr := pointer.PointerOf("hello world")
	fmt.Printf("String: %s\n", *strPtr)

	// Useful for struct fields
	type Config struct {
		Retries *int
		Tag     *string
	}

	cfg := Config{
		Retries: pointer.PointerOf(3),
		Tag:     pointer.PointerOf("v1"),
	}

	fmt.Printf("Config: Retries=%d, Tag=%s\n", *cfg.Retries, *cfg.Tag)

	// Output:
	// Int: 42
	// String: hello world
	// Config: Retries=3, Tag=v1
}
