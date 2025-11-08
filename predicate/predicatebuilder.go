package main

import "fmt"

// PredicateBuilder provides a fluent interface for building complex predicates
// This combines the Predicate pattern with the Builder pattern

// Process represents a system process
type Process struct {
	ID       int
	Title    string
	Status   string
	Priority int
	Owner    string
	CPUUsage float64
	Memory   int64
}

func (p *Process) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s, Status: %s, Priority: %d, Owner: %s}",
		p.ID, p.Title, p.Status, p.Priority, p.Owner)
}

// ProcessPredicate is a function that tests a Process
type ProcessPredicate func(*Process) bool

// ProcessPredicateBuilder builds complex process predicates
type ProcessPredicateBuilder struct {
	predicates []ProcessPredicate
	combineOp  string // "AND" or "OR"
}

// NewProcessPredicateBuilder creates a new predicate builder
func NewProcessPredicateBuilder() *ProcessPredicateBuilder {
	return &ProcessPredicateBuilder{
		predicates: []ProcessPredicate{},
		combineOp:  "AND",
	}
}

// WithTitle adds a title filter
func (b *ProcessPredicateBuilder) WithTitle(title string) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByTitle(title))
	return b
}

// WithID adds an ID filter
func (b *ProcessPredicateBuilder) WithID(id int) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByID(id))
	return b
}

// WithStatus adds a status filter
func (b *ProcessPredicateBuilder) WithStatus(status string) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByStatus(status))
	return b
}

// WithMinPriority adds a minimum priority filter
func (b *ProcessPredicateBuilder) WithMinPriority(priority int) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByMinPriority(priority))
	return b
}

// WithOwner adds an owner filter
func (b *ProcessPredicateBuilder) WithOwner(owner string) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByOwner(owner))
	return b
}

// WithMaxCPU adds a maximum CPU usage filter
func (b *ProcessPredicateBuilder) WithMaxCPU(maxCPU float64) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByMaxCPU(maxCPU))
	return b
}

// WithMinMemory adds a minimum memory filter
func (b *ProcessPredicateBuilder) WithMinMemory(minMemory int64) *ProcessPredicateBuilder {
	b.predicates = append(b.predicates, ByMinMemory(minMemory))
	return b
}

// UseAND sets the combinator to AND (default)
func (b *ProcessPredicateBuilder) UseAND() *ProcessPredicateBuilder {
	b.combineOp = "AND"
	return b
}

// UseOR sets the combinator to OR
func (b *ProcessPredicateBuilder) UseOR() *ProcessPredicateBuilder {
	b.combineOp = "OR"
	return b
}

// Build creates the final predicate
func (b *ProcessPredicateBuilder) Build() ProcessPredicate {
	if len(b.predicates) == 0 {
		return func(*Process) bool { return true }
	}

	if len(b.predicates) == 1 {
		return b.predicates[0]
	}

	if b.combineOp == "OR" {
		return func(p *Process) bool {
			for _, pred := range b.predicates {
				if pred(p) {
					return true
				}
			}
			return false
		}
	}

	// Default: AND
	return func(p *Process) bool {
		for _, pred := range b.predicates {
			if !pred(p) {
				return false
			}
		}
		return true
	}
}

// Individual predicate constructors for Process

func ByTitle(title string) ProcessPredicate {
	return func(p *Process) bool {
		return p.Title == title
	}
}

func ByID(id int) ProcessPredicate {
	return func(p *Process) bool {
		return p.ID == id
	}
}

func ByStatus(status string) ProcessPredicate {
	return func(p *Process) bool {
		return p.Status == status
	}
}

func ByMinPriority(minPriority int) ProcessPredicate {
	return func(p *Process) bool {
		return p.Priority >= minPriority
	}
}

func ByOwner(owner string) ProcessPredicate {
	return func(p *Process) bool {
		return p.Owner == owner
	}
}

func ByMaxCPU(maxCPU float64) ProcessPredicate {
	return func(p *Process) bool {
		return p.CPUUsage <= maxCPU
	}
}

func ByMinMemory(minMemory int64) ProcessPredicate {
	return func(p *Process) bool {
		return p.Memory >= minMemory
	}
}

// ProcessManager manages a collection of processes
type ProcessManager struct {
	processes []*Process
}

// CreateProcessManager creates a new process manager with sample data
func CreateProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: []*Process{
			{ID: 1, Title: "Go", Status: "running", Priority: 5, Owner: "user1", CPUUsage: 25.5, Memory: 1024},
			{ID: 2, Title: "Python", Status: "running", Priority: 3, Owner: "user2", CPUUsage: 15.2, Memory: 2048},
			{ID: 3, Title: "C++", Status: "stopped", Priority: 7, Owner: "user1", CPUUsage: 0.0, Memory: 512},
			{ID: 4, Title: "Java", Status: "running", Priority: 4, Owner: "user3", CPUUsage: 45.8, Memory: 4096},
			{ID: 5, Title: "Rust", Status: "running", Priority: 6, Owner: "user1", CPUUsage: 10.3, Memory: 1536},
			{ID: 6, Title: "Node", Status: "stopped", Priority: 2, Owner: "user2", CPUUsage: 0.0, Memory: 768},
		},
	}
}

// Find filters processes using a predicate
func (pm *ProcessManager) Find(predicate ProcessPredicate) []*Process {
	var result []*Process
	for _, p := range pm.processes {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

// GetAll returns all processes
func (pm *ProcessManager) GetAll() []*Process {
	return pm.processes
}

// Demo function showing predicate builder usage
func DemoPredicateBuilder() {
	fmt.Println("\n=== Predicate Builder Pattern Examples ===\n")

	pm := CreateProcessManager()

	// Example 1: Simple filter with builder
	fmt.Println("1. Find process by title 'Go':")
	pred1 := NewProcessPredicateBuilder().
		WithTitle("Go").
		Build()
	result1 := pm.Find(pred1)
	fmt.Printf("   Found %d process(es)\n", len(result1))
	for _, p := range result1 {
		fmt.Printf("   - %s\n", p)
	}
	fmt.Println()

	// Example 2: Multiple AND conditions
	fmt.Println("2. Find running processes with priority >= 5:")
	pred2 := NewProcessPredicateBuilder().
		WithStatus("running").
		WithMinPriority(5).
		Build()
	result2 := pm.Find(pred2)
	fmt.Printf("   Found %d process(es)\n", len(result2))
	for _, p := range result2 {
		fmt.Printf("   - %s\n", p)
	}
	fmt.Println()

	// Example 3: Owner filter
	fmt.Println("3. Find all processes owned by user1:")
	pred3 := NewProcessPredicateBuilder().
		WithOwner("user1").
		Build()
	result3 := pm.Find(pred3)
	fmt.Printf("   Found %d process(es)\n", len(result3))
	for _, p := range result3 {
		fmt.Printf("   - %s\n", p)
	}
	fmt.Println()

	// Example 4: Complex AND conditions
	fmt.Println("4. Find running processes by user1 with priority >= 5:")
	pred4 := NewProcessPredicateBuilder().
		WithStatus("running").
		WithOwner("user1").
		WithMinPriority(5).
		Build()
	result4 := pm.Find(pred4)
	fmt.Printf("   Found %d process(es)\n", len(result4))
	for _, p := range result4 {
		fmt.Printf("   - %s\n", p)
	}
	fmt.Println()

	// Example 5: Using OR combinator
	fmt.Println("5. Find processes with title 'Go' OR 'Python' (using OR):")
	pred5 := NewProcessPredicateBuilder().
		UseOR().
		WithTitle("Go").
		WithTitle("Python").
		Build()
	result5 := pm.Find(pred5)
	fmt.Printf("   Found %d process(es)\n", len(result5))
	for _, p := range result5 {
		fmt.Printf("   - %s\n", p)
	}
	fmt.Println()

	// Example 6: Resource-based filtering
	fmt.Println("6. Find processes with CPU <= 20% and Memory >= 1GB:")
	pred6 := NewProcessPredicateBuilder().
		WithMaxCPU(20.0).
		WithMinMemory(1024).
		Build()
	result6 := pm.Find(pred6)
	fmt.Printf("   Found %d process(es)\n", len(result6))
	for _, p := range result6 {
		fmt.Printf("   - %s (CPU: %.1f%%, Mem: %dMB)\n", p.Title, p.CPUUsage, p.Memory)
	}
	fmt.Println()

	// Example 7: Combining multiple criteria
	fmt.Println("7. Find high-priority running processes with low CPU:")
	pred7 := NewProcessPredicateBuilder().
		WithStatus("running").
		WithMinPriority(5).
		WithMaxCPU(30.0).
		Build()
	result7 := pm.Find(pred7)
	fmt.Printf("   Found %d process(es)\n", len(result7))
	for _, p := range result7 {
		fmt.Printf("   - %s (Priority: %d, CPU: %.1f%%)\n", p.Title, p.Priority, p.CPUUsage)
	}
	fmt.Println()

	// Example 8: Find stopped processes
	fmt.Println("8. Find all stopped processes:")
	pred8 := NewProcessPredicateBuilder().
		WithStatus("stopped").
		Build()
	result8 := pm.Find(pred8)
	fmt.Printf("   Found %d process(es)\n", len(result8))
	for _, p := range result8 {
		fmt.Printf("   - %s\n", p)
	}
}

// Advanced: Specification pattern (similar to predicate but with additional methods)

// ProcessSpecification is an interface for process specifications
type ProcessSpecification interface {
	IsSatisfiedBy(*Process) bool
	And(ProcessSpecification) ProcessSpecification
	Or(ProcessSpecification) ProcessSpecification
	Not() ProcessSpecification
}

// baseSpecification implements ProcessSpecification
type baseSpecification struct {
	predicate ProcessPredicate
}

func (s *baseSpecification) IsSatisfiedBy(p *Process) bool {
	return s.predicate(p)
}

func (s *baseSpecification) And(other ProcessSpecification) ProcessSpecification {
	return &andSpecification{s, other}
}

func (s *baseSpecification) Or(other ProcessSpecification) ProcessSpecification {
	return &orSpecification{s, other}
}

func (s *baseSpecification) Not() ProcessSpecification {
	return &notSpecification{s}
}

type andSpecification struct {
	left, right ProcessSpecification
}

func (s *andSpecification) IsSatisfiedBy(p *Process) bool {
	return s.left.IsSatisfiedBy(p) && s.right.IsSatisfiedBy(p)
}

func (s *andSpecification) And(other ProcessSpecification) ProcessSpecification {
	return &andSpecification{s, other}
}

func (s *andSpecification) Or(other ProcessSpecification) ProcessSpecification {
	return &orSpecification{s, other}
}

func (s *andSpecification) Not() ProcessSpecification {
	return &notSpecification{s}
}

type orSpecification struct {
	left, right ProcessSpecification
}

func (s *orSpecification) IsSatisfiedBy(p *Process) bool {
	return s.left.IsSatisfiedBy(p) || s.right.IsSatisfiedBy(p)
}

func (s *orSpecification) And(other ProcessSpecification) ProcessSpecification {
	return &andSpecification{s, other}
}

func (s *orSpecification) Or(other ProcessSpecification) ProcessSpecification {
	return &orSpecification{s, other}
}

func (s *orSpecification) Not() ProcessSpecification {
	return &notSpecification{s}
}

type notSpecification struct {
	spec ProcessSpecification
}

func (s *notSpecification) IsSatisfiedBy(p *Process) bool {
	return !s.spec.IsSatisfiedBy(p)
}

func (s *notSpecification) And(other ProcessSpecification) ProcessSpecification {
	return &andSpecification{s, other}
}

func (s *notSpecification) Or(other ProcessSpecification) ProcessSpecification {
	return &orSpecification{s, other}
}

func (s *notSpecification) Not() ProcessSpecification {
	return s.spec // Double negation
}

// Specification constructors

func RunningSpecification() ProcessSpecification {
	return &baseSpecification{
		predicate: func(p *Process) bool {
			return p.Status == "running"
		},
	}
}

func HighPrioritySpecification() ProcessSpecification {
	return &baseSpecification{
		predicate: func(p *Process) bool {
			return p.Priority >= 5
		},
	}
}

func OwnerSpecification(owner string) ProcessSpecification {
	return &baseSpecification{
		predicate: func(p *Process) bool {
			return p.Owner == owner
		},
	}
}

// Demo specification pattern
func DemoSpecificationPattern() {
	fmt.Println("\n=== Specification Pattern Example ===\n")

	pm := CreateProcessManager()

	// Build complex specification
	spec := RunningSpecification().
		And(HighPrioritySpecification()).
		And(OwnerSpecification("user1"))

	fmt.Println("Find: Running AND High Priority AND Owner='user1'")

	var result []*Process
	for _, p := range pm.GetAll() {
		if spec.IsSatisfiedBy(p) {
			result = append(result, p)
		}
	}

	fmt.Printf("Found %d process(es):\n", len(result))
	for _, p := range result {
		fmt.Printf("   - %s\n", p)
	}
}
