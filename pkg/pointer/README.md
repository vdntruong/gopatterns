# Pointer Package

## Overview

The `pointer` package provides utility functions for working with pointers in Go. Its primary feature is `PointerOf`, which returns a pointer to any value.

## Why use it?

In Go, you cannot take the address of a literal or a constant directly (e.g., `&10` or `&"string"` is invalid syntax). You typically have to create a temporary variable first.

```go
// Invalid
// ptr := &10

// Verbose
val := 10
ptr := &val
```

This package solves that minor annoyance, making code cleaner, especially when working with structs that have pointer fields (common in APIs, JSON, or database models to represent nullable fields).

## What it is

- `PointerOf[T any](v T) *T`: A generic function that takes a value of any type and returns a pointer to it.

## How it works

It accepts a value argument, which creates a copy of that value on the stack (or escapes to heap), and returns the address of that copy.

## When to use it

- **Struct Initialization**: When initializing structs with pointer fields using literals.
  ```go
  config := Config{
      Timeout: pointer.PointerOf(30),
      Name:    pointer.PointerOf("default"),
  }
  ```
- **Tests**: When defining expected values for pointer fields in test cases.
- **Optional Parameters**: When passing values to functions that accept pointers to optional arguments.

## When to avoid it

- **Large Structs**: Passing a large struct by value to `PointerOf` copies it. If performance is critical and the struct is huge, manually creating a variable and taking its address might be slightly better (though compiler optimizations often handle this).
- **Mutability**: Remember that `PointerOf` returns a pointer to a *copy* of the value passed in, not the original value if it was a variable.
  ```go
  x := 5
  p := pointer.PointerOf(x)
  *p = 10
  // x is still 5
  ```
  If you need to modify the original variable, take its address directly with `&x`.
