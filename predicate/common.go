package main

import "fmt"

// User represents a user in the system
type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Active   bool
	Role     string
	Country  string
}

// Approach 1: Multiple specific filter methods
// Problem: Need a new method for every filter combination

type UserRepository1 struct {
	users []User
}

func (r *UserRepository1) FindByAge(age int) []User {
	var result []User
	for _, u := range r.users {
		if u.Age == age {
			result = append(result, u)
		}
	}
	return result
}

func (r *UserRepository1) FindByRole(role string) []User {
	var result []User
	for _, u := range r.users {
		if u.Role == role {
			result = append(result, u)
		}
	}
	return result
}

func (r *UserRepository1) FindActiveUsers() []User {
	var result []User
	for _, u := range r.users {
		if u.Active {
			result = append(result, u)
		}
	}
	return result
}

func (r *UserRepository1) FindByAgeAndRole(age int, role string) []User {
	var result []User
	for _, u := range r.users {
		if u.Age == age && u.Role == role {
			result = append(result, u)
		}
	}
	return result
}

// Problem: Need exponential number of methods for all combinations!
// FindByAgeAndRoleAndCountry, FindByActiveAndRole, etc...

// Approach 2: Using multiple parameters with flags
// Problem: Complex parameter list, unclear null values

type UserRepository2 struct {
	users []User
}

type UserFilter struct {
	Age     *int
	Role    *string
	Active  *bool
	Country *string
}

func (r *UserRepository2) Find(filter UserFilter) []User {
	var result []User
	for _, u := range r.users {
		match := true

		if filter.Age != nil && u.Age != *filter.Age {
			match = false
		}
		if filter.Role != nil && u.Role != *filter.Role {
			match = false
		}
		if filter.Active != nil && u.Active != *filter.Active {
			match = false
		}
		if filter.Country != nil && u.Country != *filter.Country {
			match = false
		}

		if match {
			result = append(result, u)
		}
	}
	return result
}

// Problem: Ugly pointer syntax, hard to read:
// age := 25
// users := repo.Find(UserFilter{Age: &age})

// Approach 3: Hardcoded switch/case logic
// Problem: Not extensible, must modify code for new filters

type UserRepository3 struct {
	users []User
}

func (r *UserRepository3) FindByCondition(field string, value interface{}) []User {
	var result []User
	for _, u := range r.users {
		match := false

		switch field {
		case "age":
			if age, ok := value.(int); ok && u.Age == age {
				match = true
			}
		case "role":
			if role, ok := value.(string); ok && u.Role == role {
				match = true
			}
		case "active":
			if active, ok := value.(bool); ok && u.Active == active {
				match = true
			}
		case "country":
			if country, ok := value.(string); ok && u.Country == country {
				match = true
			}
		}

		if match {
			result = append(result, u)
		}
	}
	return result
}

// Problem:
// - Type unsafe (interface{})
// - Can't combine conditions
// - Must modify code for new fields

// Approach 4: Manual iteration everywhere
// Problem: Code duplication, error-prone

type UserRepository4 struct {
	users []User
}

func (r *UserRepository4) GetUsers() []User {
	return r.users
}

// Users must write filtering logic every time:
// var activeAdmins []User
// for _, u := range repo.GetUsers() {
//     if u.Active && u.Role == "admin" && u.Age > 18 {
//         activeAdmins = append(activeAdmins, u)
//     }
// }

// Problem: Repeated code, hard to test, easy to make mistakes

// Approach 5: SQL-like string queries
// Problem: String-based, no type safety, runtime errors

type UserRepository5 struct {
	users []User
}

func (r *UserRepository5) Query(query string) ([]User, error) {
	// Pseudo implementation - would need a parser
	var result []User

	// Example: "age > 25 AND role = 'admin'"
	// Problem: Parse string at runtime, no compile-time safety

	// Very simplified example
	for _, u := range r.users {
		// Would need complex parsing logic here
		result = append(result, u)
	}

	return result, nil
}

// Problem:
// - No type safety
// - Runtime errors instead of compile-time
// - Complex parsing logic needed
// - SQL injection-like vulnerabilities

// Demo function showing problems with traditional approaches
func DemoCommonApproaches() {
	fmt.Println("=== Common Filtering Approaches (Without Predicate Pattern) ===\n")

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25, Active: true, Role: "admin", Country: "USA"},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30, Active: true, Role: "user", Country: "UK"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 25, Active: false, Role: "admin", Country: "USA"},
		{ID: 4, Name: "David", Email: "david@example.com", Age: 35, Active: true, Role: "user", Country: "Canada"},
	}

	// Approach 1: Multiple methods
	fmt.Println("1. Multiple specific filter methods:")
	repo1 := &UserRepository1{users: users}
	admins := repo1.FindByRole("admin")
	fmt.Printf("   Found %d admins\n", len(admins))
	fmt.Println("   Problem: Need a method for every filter combination!")
	fmt.Println("   Example: FindByRole, FindByAge, FindByAgeAndRole, FindByAgeAndRoleAndCountry...")
	fmt.Println()

	// Approach 2: Parameter struct with pointers
	fmt.Println("2. Parameter struct with pointers:")
	repo2 := &UserRepository2{users: users}
	age := 25
	role := "admin"
	active := true
	filtered := repo2.Find(UserFilter{Age: &age, Role: &role, Active: &active})
	fmt.Printf("   Found %d users\n", len(filtered))
	fmt.Println("   Problem: Ugly pointer syntax, hard to read")
	fmt.Println()

	// Approach 3: String field with interface{} value
	fmt.Println("3. Hardcoded switch/case:")
	repo3 := &UserRepository3{users: users}
	result := repo3.FindByCondition("role", "admin")
	fmt.Printf("   Found %d users\n", len(result))
	fmt.Println("   Problem: Type unsafe, can't combine conditions, must modify code")
	fmt.Println()

	// Approach 4: Manual iteration
	fmt.Println("4. Manual iteration everywhere:")
	repo4 := &UserRepository4{users: users}
	var activeAdmins []User
	for _, u := range repo4.GetUsers() {
		if u.Active && u.Role == "admin" {
			activeAdmins = append(activeAdmins, u)
		}
	}
	fmt.Printf("   Found %d active admins (after manual iteration)\n", len(activeAdmins))
	fmt.Println("   Problem: Code duplication, error-prone, not reusable")
	fmt.Println()

	// Approach 5: String queries
	fmt.Println("5. SQL-like string queries:")
	repo5 := &UserRepository5{users: users}
	queryResult, _ := repo5.Query("age > 25 AND role = 'admin'")
	fmt.Printf("   Query would return %d users\n", len(queryResult))
	fmt.Println("   Problem: No type safety, runtime errors, complex parsing")
	fmt.Println()

	fmt.Println("=== The Predicate Pattern solves all these problems! ===")
}

// Helper function to demonstrate the problem of manual filtering
func ManualFilteringExample() {
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, Active: true, Role: "admin"},
		{ID: 2, Name: "Bob", Age: 30, Active: true, Role: "user"},
		{ID: 3, Name: "Charlie", Age: 25, Active: false, Role: "admin"},
	}

	// Without predicates - repeated code
	var activeUsers []User
	for _, u := range users {
		if u.Active {
			activeUsers = append(activeUsers, u)
		}
	}

	var activeAdmins []User
	for _, u := range users {
		if u.Active && u.Role == "admin" {
			activeAdmins = append(activeAdmins, u)
		}
	}

	var youngActiveAdmins []User
	for _, u := range users {
		if u.Active && u.Role == "admin" && u.Age < 30 {
			youngActiveAdmins = append(youngActiveAdmins, u)
		}
	}

	fmt.Printf("Active: %d, Active Admins: %d, Young Active Admins: %d\n",
		len(activeUsers), len(activeAdmins), len(youngActiveAdmins))

	fmt.Println("Problem: Look at all that duplicated code!")
}
