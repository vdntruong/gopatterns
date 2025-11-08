package main

import "fmt"

// Server represents a server configuration
type Server struct {
	Host     string
	Port     int
	Protocol string
	Timeout  int
	MaxConns int
}

// NewServer creates a new server with the given configuration
//
// Approach 1: Simple constructor with all parameters
// Problem: Hard to remember parameter order, not flexible
func NewServer(host string, port int, protocol string, timeout int, maxConns int) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		Timeout:  timeout,
		MaxConns: maxConns,
	}
}

// NewServerWithDefaults creates a new server with the given configuration
//
// Approach 2: Constructor with default values
// Problem: Still inflexible, requires all parameters
func NewServerWithDefaults(host string, port int) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		Protocol: "http",
		Timeout:  30,
		MaxConns: 100,
	}
}

// NewServerBasic creates a new server with the given configuration
//
// Approach 3: Telescoping constructors
// Problem: Need many constructors for different combinations
func NewServerBasic(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func NewServerWithProtocol(host string, port int, protocol string) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		Protocol: protocol,
	}
}

func NewServerWithTimeout(host string, port int, protocol string, timeout int) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		Timeout:  timeout,
	}
}

func NewServerFull(host string, port int, protocol string, timeout int, maxConns int) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		Timeout:  timeout,
		MaxConns: maxConns,
	}
}

// NewServerStruct creates a new server with the given configuration
//
// Approach 4: Struct initialization
// Problem: No validation, no encapsulation, verbose
func NewServerStruct() *Server {
	return &Server{
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
		Timeout:  30,
		MaxConns: 100,
	}
}

// ServerConfig represents a server configuration
//
// Approach 5: Using a config struct
// Problem: Still verbose, no validation until after creation
type ServerConfig struct {
	Host     string
	Port     int
	Protocol string
	Timeout  int
	MaxConns int
}

func NewServerFromConfig(config ServerConfig) *Server {
	return &Server{
		Host:     config.Host,
		Port:     config.Port,
		Protocol: config.Protocol,
		Timeout:  config.Timeout,
		MaxConns: config.MaxConns,
	}
}

func (s *Server) String() string {
	return fmt.Sprintf("Server{Host: %s, Port: %d, Protocol: %s, Timeout: %ds, MaxConns: %d}",
		s.Host, s.Port, s.Protocol, s.Timeout, s.MaxConns)
}

// DemoCommonApproaches demonstrates the problems with traditional approaches to creating objects
func DemoCommonApproaches() {
	fmt.Println("=== Common Constructor Approaches (Without Patterns) ===")
	fmt.Println()

	// Approach 1: All parameters - hard to read, easy to mix up
	server1 := NewServer("localhost", 8080, "https", 60, 200)
	fmt.Println("1. All parameters constructor:")
	fmt.Println(server1)
	fmt.Println("Problem: Easy to mix up parameters (is 60 timeout or maxConns?)")
	fmt.Println()

	// Approach 2: With defaults
	server2 := NewServerWithDefaults("api.example.com", 443)
	fmt.Println("2. Constructor with defaults:")
	fmt.Println(server2)
	fmt.Println("Problem: Can't customize protocol or timeout without creating new constructor")
	fmt.Println()

	// Approach 3: Telescoping - need different constructors
	server3 := NewServerWithProtocol("localhost", 3000, "ws")
	fmt.Println("3. Telescoping constructor:")
	fmt.Println(server3)
	fmt.Println("Problem: Need many constructors for different combinations")
	fmt.Println()

	// Approach 4: Direct struct initialization
	server4 := &Server{
		Host:     "db.example.com",
		Port:     5432,
		Protocol: "tcp",
	}
	fmt.Println("4. Direct struct initialization:")
	fmt.Println(server4)
	fmt.Println("Problem: No validation, easy to forget required fields")
	fmt.Println()

	// Approach 5: Config struct
	config := ServerConfig{
		Host:     "cache.example.com",
		Port:     6379,
		Protocol: "redis",
		Timeout:  10,
		MaxConns: 50,
	}
	server5 := NewServerFromConfig(config)
	fmt.Println("5. Config struct approach:")
	fmt.Println(server5)
	fmt.Println("Problem: Still verbose, no validation during construction")
	fmt.Println()

	fmt.Println("=== All these approaches have limitations that Builder pattern solves ===")
}
