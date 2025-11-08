package main

import "fmt"

func main() {
	printLine()
	fmt.Println("  PREDICATE PATTERN IN GO")
	printLine()

	// First, demonstrate common approaches and their problems
	DemoCommonApproaches()

	printLine()
	fmt.Println()

	// Show the generic predicate pattern solution
	DemoPredicatePattern()

	printLine()
	fmt.Println()

	// Generic predicates with simple types
	DemoGenericPredicates()

	printLine()
	fmt.Println()

	// Predicate Builder pattern
	DemoPredicateBuilder()

	printLine()
	fmt.Println()

	// Specification pattern variant
	DemoSpecificationPattern()

	printLine()
	fmt.Println("  DEMO COMPLETED")
	printLine()
}

func printLine() {
	fmt.Println("========================================")
}
