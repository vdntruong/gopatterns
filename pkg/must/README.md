# Must Package

## Overview

The `must` package provides utility functions to simplify error handling in Go programs by panicking on errors. This is particularly useful for initialization code, scripts, or scenarios where an error is unrecoverable or indicates a programmer error.

## Why use it?

In Go, error handling is explicit, which is great for robustness but can be verbose in certain situations.
- **Initialization**: When initializing global variables or setting up the application state (e.g., parsing a template, compiling a regex), an error often means the application cannot proceed.
- **Scripts/Tools**: In simple CLI tools, panicking on error might be the desired behavior.
- **Chaining**: It allows for cleaner one-liners when you are certain an operation should succeed.

## What it is

The package provides two main functions:
- `Must(err error)`: Panics if the provided error is not nil.
- `MustV[T any](v T, err error) T`: Returns the value `v` if `err` is nil, otherwise panics.

## How it works

It simply checks the error argument. If it's not `nil`, it calls `panic(err)`.

## When to use it

- **Application Startup**: Loading configuration, connecting to databases (if strict), compiling regular expressions.
  ```go
  var emailRegex = must.MustV(regexp.Compile("^[a-z]+@[a-z]+\\.[a-z]+$"))
  ```
- **Tests**: In test code where you want to fail immediately if a setup step fails (though `t.Fatal` is often preferred, `Must` can be useful in helper functions).
- **Quick Scripts**: When writing throwaway scripts where error handling boilerplate is unnecessary.

## When to avoid it

- **Libraries**: Do not use `Must` in library code intended for others. Always return errors to let the caller decide how to handle them.
- **Runtime Logic**: Avoid in the main request processing loop of a server. A panic could crash the entire application or goroutine (though `recover` middleware usually catches it, it's bad practice).
- **Recoverable Errors**: If there's a reasonable way to handle the error (e.g., retry, fallback, user notification), do not use `Must`.
