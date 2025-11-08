# Predicate Pattern in Go

## Overview

The Predicate pattern is a behavioral design pattern that uses functions returning boolean values to test whether objects satisfy certain conditions. Predicates are powerful for filtering collections, validating data, and implementing composable conditional logic in a type-safe, reusable way.

In Go, predicates leverage first-class functions and generics (Go 1.18+) to create flexible, composable filtering logic.

## Problem Statement

When working with collections and filtering data, traditional approaches create several problems:

### Problems with Traditional Approaches

1. **Multiple Specific Methods**: Need exponential methods for every filter combination
   ```go
   FindByAge(age int)
   FindByRole(role string)
   FindByAgeAndRole(age int, role string)
   FindByAgeAndRoleAndCountry(age int, role string, country string)
   // Combinatorial explosion!
   ```

2. **Parameter Structs with Pointers**: Ugly syntax, unclear null semantics
   ```go
   age := 25
   users := repo.Find(UserFilter{Age: &age}) // Awkward pointer syntax
   ```

3. **Type-Unsafe String Queries**: No compile-time safety
   ```go
   repo.Query("age > 25 AND role = 'admin'") // String parsing, runtime errors
   ```

4. **Manual Iteration**: Repeated code everywhere
   ```go
   var filtered []User
   for _, u := range users {
       if u.Age > 25 && u.Role == "admin" {
           filtered = append(filtered, u)
       }
   }
   // Duplicate this logic everywhere!
   ```

5. **Hardcoded Logic**: Can't extend without modifying code
   ```go
   switch field {
   case "age": // Must add case for every field
   case "role":
   }
   ```

## Solution: Predicate Pattern

The Predicate pattern provides:
- ✅ Type-safe filtering with compile-time checking
- ✅ Composable logic with AND, OR, NOT combinators
- ✅ Reusable predicates across codebase
- ✅ No code duplication
- ✅ Easy to extend without modifying existing code
- ✅ Works with Go generics for any type

## Implementation

### Core Components

1. **Predicate Type**: A function that tests a condition
   ```go
   type Predicate[T any] func(T) bool
   ```

2. **Filter Function**: Applies predicate to collection
   ```go
   func Filter[T any](items []T, predicate Predicate[T]) []T
   ```

3. **Predicate Constructors**: Create specific predicates
   ```go
   func ByAge(age int) Predicate[User] {
       return func(u User) bool {
           return u.Age == age
       }
   }
   ```

4. **Combinators**: Combine predicates logically
   ```go
   func And[T any](p1, p2 Predicate[T]) Predicate[T]
   func Or[T any](p1, p2 Predicate[T]) Predicate[T]
   func Not[T any](p Predicate[T]) Predicate[T]
   ```

### Basic Example

```go
// Define predicate type
type Predicate[T any] func(T) bool

// Generic filter function
func Filter[T any](items []T, pred Predicate[T]) []T {
    var result []T
    for _, item := range items {
        if pred(item) {
            result = append(result, item)
        }
    }
    return result
}

// Create specific predicates
func IsEven(n int) bool {
    return n%2 == 0
}

func GreaterThan(threshold int) Predicate[int] {
    return func(n int) bool {
        return n > threshold
    }
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6}
evenNumbers := Filter(numbers, Predicate[int](IsEven))
// [2, 4, 6]
```

## Usage Examples

### Simple Filtering

```go
products := []Product{...}

// Filter by category
electronics := Filter(products, ByCategory("Electronics"))

// Filter by price range
affordable := Filter(products, ByPriceRange(100, 500))

// Filter in-stock items
available := Filter(products, InStock())
```

### Combining Predicates

```go
// AND: Electronics that are in stock
availableElectronics := Filter(products,
    And(ByCategory("Electronics"), InStock()))

// OR: Furniture or cheap items
furnitureOrCheap := Filter(products,
    Or(ByCategory("Furniture"), ByMaxPrice(50)))

// NOT: Out of stock items
outOfStock := Filter(products, Not(InStock()))
```

### Complex Combinations

```go
// Multiple AND conditions
premiumElectronics := Filter(products,
    And(
        And(ByCategory("Electronics"), InStock()),
        And(ByMinRating(4.5), ByMinPrice(100)),
    ))

// Real-world: High-quality affordable products
goodDeals := Filter(products,
    And(
        ByMinRating(4.0),
        And(InStock(), ByMaxPrice(300)),
    ))
```

### Other Operations

```go
// Any: Check if any element matches
hasExpensive := Any(products, ByMinPrice(1000))

// All: Check if all elements match
allInStock := All(products, InStock())

// None: Check if no elements match
noOutOfStock := None(products, Not(InStock()))

// Find: Get first match
if product, found := Find(products, ByID(42)); found {
    fmt.Println(product.Name)
}

// Count: Count matches
count := Count(products, ByCategory("Electronics"))
```

## Advanced: Predicate Builder Pattern

Combine Predicate with Builder for fluent API:

```go
type PredicateBuilder struct {
    predicates []Predicate[Product]
    combineOp  string // "AND" or "OR"
}

func NewBuilder() *PredicateBuilder {
    return &PredicateBuilder{combineOp: "AND"}
}

func (b *PredicateBuilder) WithCategory(cat string) *PredicateBuilder {
    b.predicates = append(b.predicates, ByCategory(cat))
    return b
}

func (b *PredicateBuilder) InStock() *PredicateBuilder {
    b.predicates = append(b.predicates, InStock())
    return b
}

func (b *PredicateBuilder) Build() Predicate[Product] {
    // Combine all predicates with AND or OR
    return combinedPredicate
}

// Usage
pred := NewBuilder().
    WithCategory("Electronics").
    InStock().
    WithMinRating(4.5).
    Build()

results := Filter(products, pred)
```

## Advanced: Specification Pattern

Object-oriented variant with method chaining:

```go
type Specification interface {
    IsSatisfiedBy(Product) bool
    And(Specification) Specification
    Or(Specification) Specification
    Not() Specification
}

// Usage
spec := InStockSpec().
    And(ElectronicsSpec()).
    And(MinRatingSpec(4.5))

for _, product := range products {
    if spec.IsSatisfiedBy(product) {
        // Product matches
    }
}
```

## Key Benefits

### 1. Type Safety
```go
// Compile-time checking
pred := ByAge(25) // ✓ Type safe
// pred := ByAge("25") // ✗ Compile error
```

### 2. Reusability
```go
// Define once, use everywhere
activeUsers := ByActive(true)

users1 := Filter(allUsers, activeUsers)
users2 := Filter(adminUsers, activeUsers)
users3 := Filter(premiumUsers, activeUsers)
```

### 3. Composability
```go
// Build complex logic from simple predicates
basic := ByCategory("Electronics")
quality := ByMinRating(4.5)
affordable := ByMaxPrice(500)

// Compose them
pred := And(basic, And(quality, affordable))
```

### 4. Testability
```go
func TestByAge(t *testing.T) {
    pred := ByAge(25)

    assert.True(t, pred(User{Age: 25}))
    assert.False(t, pred(User{Age: 30}))
}
```

### 5. No Code Duplication
```go
// One Filter function for all types
evenNums := Filter(numbers, IsEven)
longWords := Filter(words, LongerThan(5))
activeUsers := Filter(users, IsActive)
```

## When to Use

### Use When:
- Filtering collections frequently
- Need composable filtering logic
- Want type-safe queries
- Have complex conditional logic
- Need reusable selection criteria
- Want to avoid code duplication

### Don't Use When:
- Simple one-time filtering (use inline loops)
- Performance is absolutely critical (predicates add function call overhead)
- Working with databases (use query language instead)
- Team unfamiliar with functional programming

## Real-World Examples

### E-commerce Product Filtering
```go
// Customer searches: "affordable laptops in stock with good ratings"
results := Filter(products,
    And(
        And(ByCategory("Laptop"), InStock()),
        And(ByMaxPrice(1000), ByMinRating(4.0)),
    ))
```

### Process Management
```go
// Find critical processes using too much CPU
critical := Filter(processes,
    And(
        And(ByStatus("running"), ByMinPriority(7)),
        ByMaxCPU(80.0),
    ))
```

### User Management
```go
// Find active admin users from specific country
admins := Filter(users,
    And(
        And(IsActive(), ByRole("admin")),
        ByCountry("USA"),
    ))
```

### File System
```go
// Find large log files modified recently
files := Filter(allFiles,
    And(
        And(HasExtension(".log"), LargerThan(100*MB)),
        ModifiedAfter(time.Now().AddDate(0, 0, -7)),
    ))
```

## Pattern Variations

### 1. Simple Predicates (No Errors)
```go
type Predicate[T any] func(T) bool
```

### 2. Predicates with Errors
```go
type Predicate[T any] func(T) (bool, error)
```

### 3. Predicates with Context
```go
type Predicate[T any] func(context.Context, T) bool
```

### 4. Named Specifications
```go
type Specification interface {
    IsSatisfiedBy(T) bool
}
```

## Best Practices

### 1. Name Predicates Clearly
```go
// Good
IsActive()
ByAge(25)
HasEmail()
InPriceRange(100, 500)

// Bad
Check()
Test()
Validate()
```

### 2. Make Predicates Pure
```go
// Good: No side effects
func IsEven(n int) bool {
    return n%2 == 0
}

// Bad: Has side effects
func IsEven(n int) bool {
    log.Printf("Checking %d", n) // Side effect!
    return n%2 == 0
}
```

### 3. Use Predicate Constructors
```go
// Good: Constructor returns predicate
func ByAge(age int) Predicate[User] {
    return func(u User) bool {
        return u.Age == age
    }
}

// Allows parameterization
pred := ByAge(25)
```

### 4. Provide Common Combinators
```go
func And[T any](p1, p2 Predicate[T]) Predicate[T]
func Or[T any](p1, p2 Predicate[T]) Predicate[T]
func Not[T any](p Predicate[T]) Predicate[T]
```

### 5. Document Predicate Behavior
```go
// ByMinRating creates a predicate that returns true
// if the product's rating is greater than or equal to minRating.
func ByMinRating(minRating float64) Predicate[Product] {
    return func(p Product) bool {
        return p.Rating >= minRating
    }
}
```

## Testing

Predicates are easy to test:

```go
func TestByCategory(t *testing.T) {
    pred := ByCategory("Electronics")

    tests := []struct{
        name string
        product Product
        want bool
    }{
        {"matches", Product{Category: "Electronics"}, true},
        {"no match", Product{Category: "Furniture"}, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := pred(tt.product)
            assert.Equal(t, tt.want, got)
        })
    }
}

func TestAndCombinator(t *testing.T) {
    pred := And(IsEven, GreaterThan(5))

    assert.True(t, pred(6))   // Even and > 5
    assert.True(t, pred(8))   // Even and > 5
    assert.False(t, pred(4))  // Even but not > 5
    assert.False(t, pred(7))  // > 5 but not even
}
```

## Performance Considerations

1. **Function Call Overhead**: Each predicate is a function call
   - Usually negligible compared to business logic
   - Consider inline loops for tight performance loops

2. **Short-Circuit Evaluation**: AND stops at first false, OR at first true
   ```go
   // Efficient: expensive check last
   And(IsActive(), HasComplexCalculation())
   ```

3. **Predicate Caching**: Reuse predicates when possible
   ```go
   // Good: Create once, reuse
   activeUsers := ByActive(true)
   Filter(users1, activeUsers)
   Filter(users2, activeUsers)

   // Less efficient: Create every time
   Filter(users1, ByActive(true))
   Filter(users2, ByActive(true))
   ```

## Comparison with Other Patterns

| Pattern | Purpose | Complexity | Type Safety |
|---------|---------|------------|-------------|
| **Predicate** | Filtering/selection | Low | ✅ High |
| **Strategy** | Algorithm selection | Medium | ✅ High |
| **Specification** | Business rules | Medium-High | ✅ High |
| **Chain of Responsibility** | Request handling | Medium | ✅ High |
| **Visitor** | Operations on structure | High | ✅ High |

## Summary

The Predicate pattern is a powerful, functional approach to filtering and selection in Go. It provides:

- **Type Safety**: Compile-time checking
- **Composability**: Build complex logic from simple parts
- **Reusability**: Define once, use everywhere
- **Testability**: Easy to unit test
- **Clean Code**: No duplication, clear intent

Use predicates whenever you need flexible, composable filtering logic. Combined with Go generics, it's a type-safe, elegant solution for collection operations.

## Further Reading

- [Specification Pattern](https://en.wikipedia.org/wiki/Specification_pattern) - Martin Fowler
- [Functional Options in Go](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) - Dave Cheney
- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics) - Official Go Documentation
