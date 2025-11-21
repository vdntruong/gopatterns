package must_test

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/vdntruong/gopatterns/pkg/must"
)

func ExampleMust() {
	// This will not panic because the error is nil
	must.Must(nil)
	fmt.Println("Success")

	// This would panic:
	// must.Must(errors.New("something went wrong"))

	// Output:
	// Success
}

func ExampleMustV() {
	// Useful for wrapping functions that return (T, error)
	atoi := func(s string) (int, error) {
		return strconv.Atoi(s)
	}

	// Returns the value directly if no error
	val := must.MustV(atoi("123"))
	fmt.Printf("Value: %d\n", val)

	// This would panic:
	// must.MustV(atoi("invalid"))

	// Output:
	// Value: 123
}

func ExampleMustV_customError() {
	// Demonstrating panic behavior (recover for example purposes)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fail := func() (int, error) {
		return 0, errors.New("critical failure")
	}

	must.MustV(fail())

	// Output:
	// Recovered: critical failure
}
