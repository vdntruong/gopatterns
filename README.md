# Go Design Patterns

A practical collection of software design patterns implemented in idiomatic Go. Each pattern includes detailed explanations, real-world examples, and demonstrations of traditional approaches vs. pattern-based solutions.

## 📚 Patterns Implemented

### Creational Patterns

- **[Builder Pattern](./builder/)** - Construct complex objects step-by-step with a fluent interface
  - Separates object construction from representation
  - Supports optional parameters and validation
  - Includes comprehensive examples with Computer configuration

- **[Option Pattern](./option/)** - Configure objects using functional options (Go-idiomatic)
  - Functional approach to optional parameters
  - Backward compatible and extensible
  - Examples with Server and Database configuration

### Behavioral Patterns

- **[Predicate Pattern](./predicate/)** - Type-safe filtering and selection using composable predicates
  - Generic implementation with Go 1.18+ generics
  - Composable with AND, OR, NOT combinators
  - Includes Predicate Builder and Specification pattern variants
  - Examples with Product filtering, Process management

## 🏗️ Repository Structure

Each pattern follows a consistent structure for easy learning:

```
pattern-name/
├── README.md           # Comprehensive documentation
├── main.go            # Executable demonstration
├── pattern.go         # Main pattern implementation
├── common.go          # Traditional approaches (anti-patterns)
└── ...                # Additional pattern variations
```

### File Descriptions

- **README.md** - Detailed explanation including:
  - Problem statement and motivation
  - Pattern solution and benefits
  - Implementation details with code examples
  - Usage examples (simple to complex)
  - When to use/avoid guidelines
  - Best practices and testing strategies
  - Real-world examples

- **main.go** - Executable demonstrations:
  - Shows traditional approaches and their problems
  - Demonstrates pattern implementation
  - Multiple examples showcasing features
  - Run with `go run .` in pattern directory

- **pattern.go** - Core pattern implementation:
  - Clean, production-ready code
  - Well-documented with comments
  - Generic implementations where applicable
  - Type-safe and idiomatic Go

- **common.go** - Traditional approaches:
  - Shows problems without the pattern
  - Multiple anti-pattern examples
  - Helps understand why the pattern is needed

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher (for generic patterns)
- Basic understanding of Go

### Running Examples

Navigate to any pattern directory and run:

```bash
cd builder
go run .
```

Or run specific pattern files:

```bash
cd option
go run main.go
```

## 📖 Learning Path

### Recommended Order for Beginners

1. **Builder Pattern** - Start here for object creation
2. **Option Pattern** - Learn Go-idiomatic configuration
3. **Predicate Pattern** - Master filtering and composition

### Key Concepts Covered

- **Fluent Interfaces** - Method chaining for readable code
- **Functional Programming** - First-class functions, closures
- **Generics** - Type-safe, reusable code (Go 1.18+)
- **Composition** - Building complex behavior from simple parts
- **SOLID Principles** - Single Responsibility, Open/Closed, etc.

## 💡 Pattern Comparison

| Pattern | Use Case | Complexity | Go-Idiomatic |
|---------|----------|------------|--------------|
| **Builder** | Complex object construction | Low-Medium | ⚠️ Less common |
| **Option** | Optional parameters | Low | ✅ Highly idiomatic |
| **Predicate** | Filtering & selection | Low-Medium | ✅ With generics |

### When to Use Each Pattern

**Builder Pattern**:
- Object has many optional parameters (4+)
- Construction is complex and multi-step
- Need immutable objects
- Want method chaining

**Option Pattern**:
- Configuring objects with many optional parameters
- Need backward compatibility
- Want idiomatic Go code
- Library/API design

**Predicate Pattern**:
- Filtering collections frequently
- Need composable conditional logic
- Want type-safe queries
- Have complex filtering requirements

## 🎯 Code Quality

Each pattern implementation includes:

- ✅ **Comprehensive documentation** - Detailed README with examples
- ✅ **Runnable demonstrations** - Working code you can execute
- ✅ **Anti-pattern examples** - Shows problems without the pattern
- ✅ **Best practices** - Idiomatic Go code
- ✅ **Real-world examples** - Practical use cases
- ✅ **Type safety** - Leverages Go's type system
- ✅ **No external dependencies** - Uses only standard library

## 📝 Contributing

Contributions are welcome! To add a new pattern:

1. Follow the existing structure (README.md, main.go, pattern.go, common.go)
2. Include comprehensive documentation
3. Add runnable examples
4. Show traditional approaches vs. pattern solution
5. Include best practices and when to use/avoid
6. Ensure code is idiomatic Go

## 📚 Additional Resources

### Go Design Patterns
- [Effective Go](https://go.dev/doc/effective_go) - Official Go best practices
- [Go Proverbs](https://go-proverbs.github.io/) - Simple, poetic, pithy programming proverbs
- [Dave Cheney's Blog](https://dave.cheney.net/) - Go patterns and best practices

### Design Patterns in General
- [Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns) - Gang of Four
- [Refactoring Guru](https://refactoring.guru/design-patterns) - Visual pattern explanations
- [Martin Fowler's Blog](https://martinfowler.com/) - Software design and patterns

## 🏷️ Tags

`go` `golang` `design-patterns` `creational-patterns` `behavioral-patterns` `builder-pattern` `option-pattern` `predicate-pattern` `functional-programming` `generics` `best-practices` `clean-code`

## 📄 License

This project is intended for educational purposes. Feel free to use the code in your projects.

---

⭐ **Star this repository** if you find it helpful!

🔔 **Watch** for updates as new patterns are added regularly.
