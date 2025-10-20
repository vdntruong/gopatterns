package main

func IsEven(n int) bool {
	return n%2 == 0
}

func HasPrefix(s string, prefix string) bool {
	return s[:len(prefix)] == prefix
}
