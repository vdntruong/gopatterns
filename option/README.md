# Functional Options Pattern in Go

## Overview

The Functional Options pattern (also known as the Options pattern or Configuration pattern) is a Go-idiomatic design pattern for configuring objects with many optional parameters. It uses variadic functions and closures to provide a clean, flexible, and extensible API for object construction.

This pattern was popularized by Rob Pike's blog post "Self-referential functions and the design of options" and Dave Cheney's work on functional options.

## Problem Statement

When creating objects with many configuration parameters, traditional approaches have significant drawbacks:

### Problems with Traditional Approaches

1. **Telescoping Constructors**: Creating multiple constructors for different parameter combinations
   ```go
   NewServer(host, port)
   NewServerWithTimeout(host, port, timeout)
   NewServerWithTimeoutAndTLS(host, port, timeout, enableTLS)
   // ... exponential growth!
   ```

2. **Too Many Parameters**: Functions with many parameters are hard to read and error-prone
   ```go
   // Which parameter is which?
   server := NewServer("localhost", 8080, 30*time.Second, 100, true, "/cert", "/key", logger, 3)
   ```

3. **Config Structs**: Verbose and no validation during construction
   ```go
   config := ServerConfig{
       Host: "localhost",
       Port: 8080,
       // Easy to forget required fields
       // No validation until used
   }
   ```

4. **Setters**: Objects can be in invalid state between construction and full configuration
   ```go
   server := NewServer()
   server.SetHost("localhost")
   // Oops! Forgot to set port, server is in invalid state
   server.Start() // May crash or behave unexpectedly
   ```

## Solution: Functional Options Pattern

The Options pattern provides:
- ✅ Readable, self-documenting code
- ✅ Sensible defaults with optional overrides
- ✅ Easy to add new options without breaking existing code
- ✅ Validation during construction
- ✅ Immutability (if fields are private)
- ✅ Composable and reusable options

## Implementation

### Core Components

1. **Option Type**: A function that modifies the object
   ```go
   type ServerOption func(*Server) error
   ```

2. **Constructor**: Accepts variadic options
   ```go
   func NewServer(opts ...ServerOption) (*Server, error)
   ```

3. **Option Functions**: Return Option closures
   ```go
   func WithPort(port int) ServerOption {
       return func(s *Server) error {
           s.port = port
           return nil
       }
   }
   ```

### Complete Example

```go
// The object being configured
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

// Option type
type ServerOption func(*Server) error

// Constructor with defaults
func NewServer(opts ...ServerOption) (*Server, error) {
    // Set defaults
    s := &Server{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
    }

    // Apply options
    for _, opt := range opts {
        if err := opt(s); err != nil {
            return nil, err
        }
    }

    // Validate
    if err := s.validate(); err != nil {
        return nil, err
    }

    return s, nil
}

// Option functions
func WithHost(host string) ServerOption {
    return func(s *Server) error {
        s.host = host
        return nil
    }
}

func WithPort(port int) ServerOption {
    return func(s *Server) error {
        if port < 1 || port > 65535 {
            return errors.New("invalid port")
        }
        s.port = port
        return nil
    }
}
```

## Usage Examples

### Basic Usage

```go
// Use defaults
server, err := NewServer()

// Override specific options
server, err := NewServer(
    WithHost("api.example.com"),
    WithPort(443),
)

// Mix and match options
server, err := NewServer(
    WithPort(8000),
    WithTimeout(60 * time.Second),
    WithTLS("/cert.pem", "/key.pem"),
)
```

### Advanced: Composable Options

Create preset configurations by combining options:

```go
// Production configuration preset
func WithProduction() ServerOption {
    return func(s *Server) error {
        opts := []ServerOption{
            WithHost("0.0.0.0"),
            WithPort(443),
            WithMaxConnections(1000),
            WithTimeout(60 * time.Second),
            WithMiddleware("logging", "recovery"),
        }

        for _, opt := range opts {
            if err := opt(s); err != nil {
                return err
            }
        }
        return nil
    }
}

// Usage
server, err := NewServer(
    WithProduction(),
    WithTLS("/cert.pem", "/key.pem"),
)
```

### With Validation

Options can include validation logic:

```go
func WithPort(port int) ServerOption {
    return func(s *Server) error {
        if port < 1 || port > 65535 {
            return fmt.Errorf("invalid port: %d", port)
        }
        s.port = port
        return nil
    }
}

// Validation happens during construction
server, err := NewServer(
    WithPort(99999), // Returns error immediately
)
if err != nil {
    log.Fatal(err) // "invalid port: 99999"
}
```

## Key Benefits

### 1. Backward Compatibility
Adding new options doesn't break existing code:

```go
// Version 1.0
server := NewServer(WithPort(8080))

// Version 2.0 - added new option
func WithCompression(enabled bool) ServerOption { ... }

// Old code still works!
server := NewServer(WithPort(8080))

// New code can use new option
server := NewServer(
    WithPort(8080),
    WithCompression(true), // New!
)
```

### 2. Self-Documenting
Options make code intention clear:

```go
// Clear what each parameter does
server := NewServer(
    WithHost("api.example.com"),
    WithPort(443),
    WithTLS("/cert.pem", "/key.pem"),
    WithMaxConnections(500),
)

// vs unclear positional parameters
server := NewServer("api.example.com", 443, "/cert.pem", "/key.pem", 500)
```

### 3. Optional Parameters
Natural support for optional configuration:

```go
// Minimal configuration
server := NewServer()

// Full configuration
server := NewServer(
    WithHost("example.com"),
    WithPort(443),
    WithTLS("/cert", "/key"),
    WithTimeout(30 * time.Second),
    WithLogger(customLogger),
    WithMiddleware("cors", "auth"),
)
```

### 4. Immutability
Private fields with getters prevent external modification:

```go
type Server struct {
    host string  // private
    port int     // private
}

func (s *Server) Host() string { return s.host }
func (s *Server) Port() int { return s.port }

// Configuration locked after construction
server, _ := NewServer(WithPort(8080))
// server.port = 9000  // Compile error - field is private
```

## Pattern Variations

### 1. Options Without Error Returns

Simpler when validation isn't needed:

```go
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

### 2. Options with Multiple Parameters

Bundle related configuration:

```go
func WithTLS(certFile, keyFile string) ServerOption {
    return func(s *Server) error {
        s.enableTLS = true
        s.certFile = certFile
        s.keyFile = keyFile
        return nil
    }
}
```

### 3. Options with Variadic Parameters

For list-like configuration:

```go
func WithMiddleware(middleware ...string) ServerOption {
    return func(s *Server) error {
        s.middleware = append(s.middleware, middleware...)
        return nil
    }
}

// Usage
server := NewServer(
    WithMiddleware("cors", "auth", "logging"),
)
```

## When to Use

### Use When:
- Object has 3+ optional configuration parameters
- You want sensible defaults with optional overrides
- Configuration may grow over time
- You want backward compatibility
- Validation during construction is important
- You need composable/reusable configurations

### Don't Use When:
- Object has only 1-2 simple parameters
- All parameters are required
- Performance is absolutely critical (minimal overhead but exists)
- Team unfamiliar with functional programming concepts

## Real-World Examples

### gRPC
```go
conn, err := grpc.Dial(
    "localhost:50051",
    grpc.WithInsecure(),
    grpc.WithTimeout(10*time.Second),
    grpc.WithBlock(),
)
```

### Zap Logger
```go
logger, _ := zap.NewProduction(
    zap.AddCaller(),
    zap.AddStacktrace(zapcore.ErrorLevel),
    zap.Fields(zap.String("service", "my-app")),
)
```

### HTTP Server
```go
srv := &http.Server{
    Addr:         ":8080",
    Handler:      handler,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
}
```

Could be improved with options:
```go
srv, err := NewHTTPServer(
    WithAddr(":8080"),
    WithHandler(handler),
    WithReadTimeout(15*time.Second),
    WithWriteTimeout(15*time.Second),
)
```

## Best Practices

### 1. Provide Sensible Defaults
```go
func NewServer(opts ...ServerOption) (*Server, error) {
    s := &Server{
        host:    "localhost",  // Good default
        port:    8080,         // Good default
        timeout: 30 * time.Second,
    }
    // ...
}
```

### 2. Validate Early
```go
func NewServer(opts ...ServerOption) (*Server, error) {
    s := &Server{...}

    for _, opt := range opts {
        if err := opt(s); err != nil {
            return nil, err  // Fail fast
        }
    }

    if err := s.validate(); err != nil {
        return nil, err
    }

    return s, nil
}
```

### 3. Use Private Fields
```go
type Server struct {
    host string  // private - use getter
    port int     // private - use getter
}

func (s *Server) Host() string { return s.host }
func (s *Server) Port() int { return s.port }
```

### 4. Name Options Clearly
```go
// Good
WithPort(8080)
WithTimeout(30 * time.Second)
WithTLS("/cert", "/key")

// Bad
Port(8080)
Timeout(30 * time.Second)
TLS("/cert", "/key")
```

### 5. Document Options
```go
// WithPort sets the server port.
// Port must be between 1 and 65535.
func WithPort(port int) ServerOption {
    return func(s *Server) error {
        if port < 1 || port > 65535 {
            return errors.New("port must be between 1 and 65535")
        }
        s.port = port
        return nil
    }
}
```

## Testing Benefits

Options make testing easier:

```go
func TestServerStart(t *testing.T) {
    // Easy to create test configurations
    server, err := NewServer(
        WithHost("localhost"),
        WithPort(0), // Random available port
        WithTimeout(1 * time.Second),
    )
    require.NoError(t, err)

    err = server.Start()
    assert.NoError(t, err)
}

func TestServerWithTLS(t *testing.T) {
    server, err := NewServer(
        WithTLS("testdata/cert.pem", "testdata/key.pem"),
    )
    require.NoError(t, err)
    assert.True(t, server.EnableTLS())
}
```

## Comparison with Builder Pattern

| Aspect | Options Pattern | Builder Pattern |
|--------|----------------|-----------------|
| **Style** | Functional | Object-oriented |
| **Syntax** | `NewServer(WithPort(8080))` | `NewBuilder().SetPort(8080).Build()` |
| **Immutability** | Natural (private fields) | Requires discipline |
| **Validation** | During construction | In Build() method |
| **Go Idiomatic** | ✅ Yes | ❌ Less common in Go |
| **Verbosity** | More concise | More verbose |
| **Method Chaining** | No | Yes |

## Summary

The Functional Options pattern is the Go-idiomatic way to handle configuration with many optional parameters. It provides:

- **Clean API**: Self-documenting, readable code
- **Flexibility**: Easy to add options without breaking changes
- **Safety**: Validation during construction, immutability
- **Composability**: Reusable preset configurations
- **Defaults**: Sensible defaults with selective overrides

Use this pattern whenever you have objects with multiple optional configuration parameters. It's widely used in the Go ecosystem and considered a best practice for library design.

## Further Reading

- [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html) - Rob Pike
- [Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) - Dave Cheney
- [Go Best Practices: Options Pattern](https://golang.cafe/blog/golang-functional-options-pattern.html)
