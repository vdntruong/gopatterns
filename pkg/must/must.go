package must

// Must panics if the error is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// MustV returns the value or panics if the error is not nil.
func MustV[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
