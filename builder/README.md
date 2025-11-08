# Builder Pattern in Go

## Overview

The Builder pattern is a creational design pattern that provides a flexible solution for constructing complex objects. It separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

## Problem Statement

When creating objects with many optional parameters or configurations, you typically face these challenges:

1. **Telescoping Constructor Anti-pattern**: Creating multiple constructors with different parameter combinations becomes unwieldy
2. **Unclear Object State**: Hard to understand what configuration an object has without examining all parameters
3. **Immutability Concerns**: Difficult to create immutable objects with many fields
4. **Readability**: Code becomes hard to read when constructing objects with many parameters

### Example of the Problem

```go
// Bad: Telescoping constructors
func NewComputer(cpu string) *Computer { ... }
func NewComputerWithRAM(cpu string, ram int) *Computer { ... }
func NewComputerWithRAMAndStorage(cpu string, ram int, storage int) *Computer { ... }
// ... and so on

// Bad: Too many parameters, unclear what each means
computer := NewComputer("Intel i9", 32, 1000, "RTX 4090", "Windows 11")
```

## Solution: Builder Pattern

The Builder pattern solves these problems by:

1. Providing a step-by-step interface for object construction
2. Supporting method chaining for fluent, readable code
3. Allowing optional parameters without telescoping constructors
4. Separating construction logic from the object itself

## Implementation

### Core Components

1. **Product (Computer)**: The complex object being constructed
2. **Builder (ComputerBuilder)**: Provides methods to configure and build the product
3. **Constructor (NewComputerBuilder)**: Creates a new builder instance
4. **Setter Methods**: Configure individual aspects of the product
5. **Build Method**: Returns the final constructed product

### Class Diagram

```
┌─────────────────────┐
│      Computer       │
├─────────────────────┤
│ - CPU: string       │
│ - RAM: int          │
│ - Storage: int      │
│ - GPU: string       │
│ - OS: string        │
└─────────────────────┘
         ▲
         │ creates
         │
┌─────────────────────────┐
│   ComputerBuilder       │
├─────────────────────────┤
│ - computer: *Computer   │
├─────────────────────────┤
│ + SetCPU()             │
│ + SetRAM()             │
│ + SetStorage()         │
│ + SetGPU()             │
│ + SetOS()              │
│ + Build()              │
└─────────────────────────┘
```

## Usage Examples

### Basic Usage

```go
// Build a gaming computer with all specifications
gamingPC := NewComputerBuilder().
    SetCPU("Intel Core i9").
    SetRAM(32).
    SetStorage(1000).
    SetGPU("NVIDIA RTX 4090").
    SetOS("Windows 11").
    Build()
```

### Partial Configuration

```go
// Build a basic office computer (GPU optional)
officePC := NewComputerBuilder().
    SetCPU("Intel Core i5").
    SetRAM(16).
    SetStorage(512).
    SetOS("Windows 11").
    Build()
```

### Programmatic Building

```go
builder := NewComputerBuilder()
builder.SetCPU("AMD Ryzen 9")
builder.SetRAM(64)

if needsGaming {
    builder.SetGPU("NVIDIA RTX 4090")
    builder.SetStorage(2000)
} else {
    builder.SetStorage(512)
}

computer := builder.Build()
```

## Key Benefits

### 1. Readability
The fluent interface makes code self-documenting:
```go
// Clear and readable
pc := NewComputerBuilder().
    SetCPU("Intel i9").
    SetRAM(32).
    Build()

// vs unclear constructor
pc := NewComputer("Intel i9", 32, 0, "", "")
```

### 2. Flexibility
Optional parameters are natural:
```go
// Only set what you need
basicPC := NewComputerBuilder().
    SetCPU("Intel i5").
    SetRAM(8).
    Build()
```

### 3. Maintainability
Adding new fields doesn't break existing code:
```go
// Easy to add new SetMotherboard() method without breaking existing builders
```

### 4. Validation
Can add validation in Build() method:
```go
func (b *ComputerBuilder) Build() (*Computer, error) {
    if b.computer.CPU == "" {
        return nil, errors.New("CPU is required")
    }
    return b.computer, nil
}
```

## When to Use the Builder Pattern

### Use When:
- Object has many parameters (typically 4+)
- Object has optional parameters
- Object construction is complex and multi-step
- You want immutable objects
- You need different representations of an object

### Don't Use When:
- Object is simple with few parameters
- All parameters are required
- Object construction is straightforward
- Performance is critical (builder adds slight overhead)

## Variations

### 1. Director Pattern
Add a Director to encapsulate common construction sequences:

```go
type Director struct{}

func (d *Director) BuildGamingPC() *Computer {
    return NewComputerBuilder().
        SetCPU("Intel Core i9").
        SetRAM(32).
        SetStorage(1000).
        SetGPU("NVIDIA RTX 4090").
        SetOS("Windows 11").
        Build()
}

func (d *Director) BuildOfficePC() *Computer {
    return NewComputerBuilder().
        SetCPU("Intel Core i5").
        SetRAM(16).
        SetStorage(512).
        SetOS("Windows 11").
        Build()
}
```

### 2. Functional Options Pattern (Go Idiom)
Alternative Go-idiomatic approach:

```go
type Option func(*Computer)

func WithCPU(cpu string) Option {
    return func(c *Computer) {
        c.CPU = cpu
    }
}

func NewComputer(opts ...Option) *Computer {
    c := &Computer{}
    for _, opt := range opts {
        opt(c)
    }
    return c
}

// Usage
pc := NewComputer(
    WithCPU("Intel i9"),
    WithRAM(32),
)
```

## Real-World Examples

### HTTP Request Builder
```go
request := NewHTTPRequestBuilder().
    SetURL("https://api.example.com").
    SetMethod("POST").
    AddHeader("Content-Type", "application/json").
    SetBody(jsonData).
    SetTimeout(30 * time.Second).
    Build()
```

### SQL Query Builder
```go
query := NewQueryBuilder().
    Select("id", "name", "email").
    From("users").
    Where("age > ?", 18).
    OrderBy("name ASC").
    Limit(10).
    Build()
```

### Configuration Builder
```go
config := NewConfigBuilder().
    SetHost("localhost").
    SetPort(8080).
    EnableSSL(true).
    SetTimeout(30 * time.Second).
    AddMiddleware(loggingMiddleware).
    Build()
```

## Testing Benefits

Builders make tests more readable:

```go
func TestProcessComputer(t *testing.T) {
    // Arrange
    testComputer := NewComputerBuilder().
        SetCPU("Test CPU").
        SetRAM(16).
        Build()

    // Act
    result := ProcessComputer(testComputer)

    // Assert
    assert.NotNil(t, result)
}
```

## Comparison with Other Patterns

| Pattern | Purpose | When to Use |
|---------|---------|-------------|
| **Builder** | Construct complex objects step-by-step | Many optional parameters, complex construction |
| **Factory** | Create objects without specifying exact class | Simple creation, need polymorphism |
| **Prototype** | Clone existing objects | Need copies of objects, expensive initialization |
| **Abstract Factory** | Create families of related objects | Need consistent object families |

## Summary

The Builder pattern is ideal for constructing complex objects in Go. It provides:
- Clean, readable code through method chaining
- Flexibility with optional parameters
- Separation of construction from representation
- Easy maintenance and extensibility

Use it when your objects have multiple configuration options, and you want your code to be self-documenting and maintainable.
