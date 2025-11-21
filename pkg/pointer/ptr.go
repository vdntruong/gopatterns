package pointer

// PointerOf returns a pointer to the input value.
func PointerOf[T any](v T) *T {
	return &v
}
