package main

import (
	"fmt"
	"strings"
)

// Predicate is a generic function that tests a condition on type T
type Predicate[T any] func(T) bool

// Filter applies a predicate to filter a slice
func Filter[T any](items []T, predicate Predicate[T]) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Any returns true if at least one element satisfies the predicate
func Any[T any](items []T, predicate Predicate[T]) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate
func All[T any](items []T, predicate Predicate[T]) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// None returns true if no elements satisfy the predicate
func None[T any](items []T, predicate Predicate[T]) bool {
	return !Any(items, predicate)
}

// Find returns the first element that satisfies the predicate
func Find[T any](items []T, predicate Predicate[T]) (T, bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// Count returns the number of elements that satisfy the predicate
func Count[T any](items []T, predicate Predicate[T]) int {
	count := 0
	for _, item := range items {
		if predicate(item) {
			count++
		}
	}
	return count
}

// Predicate combinators

// And combines two predicates with logical AND
func And[T any](p1, p2 Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return p1(item) && p2(item)
	}
}

// Or combines two predicates with logical OR
func Or[T any](p1, p2 Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return p1(item) || p2(item)
	}
}

// Not negates a predicate
func Not[T any](p Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return !p(item)
	}
}

// Product represents a product for demonstration
type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	InStock  bool
	Rating   float64
	Tags     []string
	Supplier string
}

func (p Product) String() string {
	return fmt.Sprintf("{ID: %d, Name: %s, Price: $%.2f, InStock: %v, Rating: %.1f}",
		p.ID, p.Name, p.Price, p.InStock, p.Rating)
}

// Product predicates

// ByCategory creates a predicate that filters by category
func ByCategory(category string) Predicate[Product] {
	return func(p Product) bool {
		return p.Category == category
	}
}

// ByPriceRange creates a predicate that filters by price range
func ByPriceRange(min, max float64) Predicate[Product] {
	return func(p Product) bool {
		return p.Price >= min && p.Price <= max
	}
}

// InStock creates a predicate for in-stock products
func InStock() Predicate[Product] {
	return func(p Product) bool {
		return p.InStock
	}
}

// ByMinRating creates a predicate for minimum rating
func ByMinRating(minRating float64) Predicate[Product] {
	return func(p Product) bool {
		return p.Rating >= minRating
	}
}

// ByNameContains creates a predicate for name search
func ByNameContains(substring string) Predicate[Product] {
	return func(p Product) bool {
		return strings.Contains(strings.ToLower(p.Name), strings.ToLower(substring))
	}
}

// BySupplier creates a predicate for filtering by supplier
func BySupplier(supplier string) Predicate[Product] {
	return func(p Product) bool {
		return p.Supplier == supplier
	}
}

// HasTag creates a predicate for checking if product has a tag
func HasTag(tag string) Predicate[Product] {
	return func(p Product) bool {
		for _, t := range p.Tags {
			if t == tag {
				return true
			}
		}
		return false
	}
}

// ByMaxPrice creates a predicate for maximum price
func ByMaxPrice(maxPrice float64) Predicate[Product] {
	return func(p Product) bool {
		return p.Price <= maxPrice
	}
}

// ByMinPrice creates a predicate for minimum price
func ByMinPrice(minPrice float64) Predicate[Product] {
	return func(p Product) bool {
		return p.Price >= minPrice
	}
}

// DemoPredicatePattern shows how to use predicates to filter a collection of products
// Demo function showing predicate pattern usage
func DemoPredicatePattern() {
	fmt.Println("=== Predicate Pattern Examples ===")

	products := []Product{
		{ID: 1, Name: "Laptop", Category: "Electronics", Price: 999.99, InStock: true, Rating: 4.5, Supplier: "TechCorp", Tags: []string{"computer", "portable"}},
		{ID: 2, Name: "Mouse", Category: "Electronics", Price: 29.99, InStock: true, Rating: 4.2, Supplier: "TechCorp", Tags: []string{"accessory", "wireless"}},
		{ID: 3, Name: "Desk", Category: "Furniture", Price: 299.99, InStock: false, Rating: 4.0, Supplier: "FurnitureCo", Tags: []string{"office", "wooden"}},
		{ID: 4, Name: "Chair", Category: "Furniture", Price: 199.99, InStock: true, Rating: 4.7, Supplier: "FurnitureCo", Tags: []string{"office", "ergonomic"}},
		{ID: 5, Name: "Monitor", Category: "Electronics", Price: 399.99, InStock: true, Rating: 4.6, Supplier: "TechCorp", Tags: []string{"display", "4k"}},
		{ID: 6, Name: "Keyboard", Category: "Electronics", Price: 79.99, InStock: false, Rating: 4.3, Supplier: "TechCorp", Tags: []string{"accessory", "mechanical"}},
	}

	// Example 1: Simple filtering
	fmt.Println("1. Filter by category:")
	electronics := Filter(products, ByCategory("Electronics"))
	fmt.Printf("   Found %d electronics\n", len(electronics))
	for _, p := range electronics {
		fmt.Printf("   - %s", p.Name)
	}
	fmt.Println()

	// Example 2: Filter by price range
	fmt.Println("2. Filter by price range ($100-$400):")
	midRange := Filter(products, ByPriceRange(100, 400))
	fmt.Printf("   Found %d products", len(midRange))
	for _, p := range midRange {
		fmt.Printf("   - %s: $%.2f\n", p.Name, p.Price)
	}
	fmt.Println()

	// Example 3: Combining predicates with AND
	fmt.Println("3. Combine predicates (Electronics AND InStock):")
	availableElectronics := Filter(products, And(ByCategory("Electronics"), InStock()))
	fmt.Printf("   Found %d available electronics\n", len(availableElectronics))
	for _, p := range availableElectronics {
		fmt.Printf("   - %s", p.Name)
	}
	fmt.Println()

	// Example 4: Complex combinations
	fmt.Println("4. Complex combination (InStock AND Price<$300 AND Rating>=4.5):")
	affordableQuality := Filter(products,
		And(
			And(InStock(), ByMaxPrice(300)),
			ByMinRating(4.5),
		),
	)
	fmt.Printf("   Found %d products", len(affordableQuality))
	for _, p := range affordableQuality {
		fmt.Printf("   - %s: $%.2f (Rating: %.1f)\n", p.Name, p.Price, p.Rating)
	}
	fmt.Println()

	// Example 5: Using OR
	fmt.Println("5. Using OR (Category=Furniture OR Price<$50):")
	furnitureOrCheap := Filter(products,
		Or(ByCategory("Furniture"), ByMaxPrice(50)),
	)
	fmt.Printf("   Found %d products\n", len(furnitureOrCheap))
	for _, p := range furnitureOrCheap {
		fmt.Printf("   - %s: $%.2f\n", p.Name, p.Price)
	}
	fmt.Println()

	// Example 6: Using NOT
	fmt.Println("6. Using NOT (NOT InStock):")
	outOfStock := Filter(products, Not(InStock()))
	fmt.Printf("   Found %d out-of-stock products\n", len(outOfStock))
	for _, p := range outOfStock {
		fmt.Printf("   - %s\n", p.Name)
	}
	fmt.Println()

	// Example 7: Name search
	fmt.Println("7. Search by name (contains 'key'):")
	matchingName := Filter(products, ByNameContains("key"))
	fmt.Printf("   Found %d products\n", len(matchingName))
	for _, p := range matchingName {
		fmt.Printf("   - %s\n", p.Name)
	}
	fmt.Println()

	// Example 8: Using Any
	fmt.Println("8. Check if any product is over $500:")
	hasExpensive := Any(products, ByMinPrice(500))
	fmt.Printf("   Has expensive products: %v\n", hasExpensive)
	fmt.Println()

	// Example 9: Using All
	fmt.Println("9. Check if all products have rating >= 4.0:")
	allHighRated := All(products, ByMinRating(4.0))
	fmt.Printf("   All products highly rated: %v\n", allHighRated)
	fmt.Println()

	// Example 10: Count
	fmt.Println("10. Count products from TechCorp:")
	techCorpCount := Count(products, BySupplier("TechCorp"))
	fmt.Printf("    TechCorp products: %d\n", techCorpCount)
	fmt.Println()

	// Example 11: Find first
	fmt.Println("11. Find first furniture item:")
	if furniture, found := Find(products, ByCategory("Furniture")); found {
		fmt.Printf("    Found: %s\n", furniture.Name)
	}
	fmt.Println()

	// Example 12: Complex real-world scenario
	fmt.Println("12. Real-world scenario: Premium in-stock electronics")
	fmt.Println("    (Electronics AND InStock AND Rating>=4.5 AND Price>=100)")
	premiumElectronics := Filter(products,
		And(
			And(ByCategory("Electronics"), InStock()),
			And(ByMinRating(4.5), ByMinPrice(100)),
		),
	)
	fmt.Printf("    Found %d premium electronics\n", len(premiumElectronics))
	for _, p := range premiumElectronics {
		fmt.Printf("    - %s: $%.2f (Rating: %.1f)\n", p.Name, p.Price, p.Rating)
	}
}

// Generic predicate examples for different types

// Integer predicates

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return n%2 != 0
}

func IsPositive(n int) bool {
	return n > 0
}

func GreaterThan(threshold int) Predicate[int] {
	return func(n int) bool {
		return n > threshold
	}
}

func LessThan(threshold int) Predicate[int] {
	return func(n int) bool {
		return n < threshold
	}
}

func Between(min, max int) Predicate[int] {
	return func(n int) bool {
		return n >= min && n <= max
	}
}

// String predicates

func HasPrefix(prefix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}

func HasSuffix(suffix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}
}

func Contains(substring string) Predicate[string] {
	return func(s string) bool {
		return strings.Contains(s, substring)
	}
}

func LongerThan(length int) Predicate[string] {
	return func(s string) bool {
		return len(s) > length
	}
}

// DemoGenericPredicates shows how to use generic predicates to filter a collection of items
func DemoGenericPredicates() {
	fmt.Println("=== Generic Predicate Examples ===")

	// Integer filtering
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("1. Filter even numbers:")
	evenNumbers := Filter(numbers, Predicate[int](IsEven))
	fmt.Printf("   %v\n", evenNumbers)

	fmt.Println("2. Filter numbers > 5:")
	greaterThan5 := Filter(numbers, GreaterThan(5))
	fmt.Printf("   %v\n", greaterThan5)

	fmt.Println("3. Filter numbers between 3 and 7:")
	between3and7 := Filter(numbers, Between(3, 7))
	fmt.Printf("   %v\n", between3and7)

	// String filtering
	words := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}

	fmt.Println("4. Filter words starting with 'b':")
	startsWithB := Filter(words, HasPrefix("b"))
	fmt.Printf("   %v\n\n", startsWithB)

	fmt.Println("5. Filter words longer than 5 characters:")
	longWords := Filter(words, LongerThan(5))
	fmt.Printf("   %v\n\n", longWords)

	fmt.Println("6. Complex: Even numbers greater than 5:")
	evenAndGreater := Filter(numbers, And(Predicate[int](IsEven), GreaterThan(5)))
	fmt.Printf("   %v\n", evenAndGreater)
}
