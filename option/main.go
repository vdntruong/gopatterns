package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Client represents the original simple example
type Client struct {
	Logger  *log.Logger
	Timeout time.Duration
	Retries int
}

func main() {
	printLine()
	fmt.Println("  FUNCTIONAL OPTIONS PATTERN IN GO")
	printLine()

	// First, demonstrate common approaches and their problems
	DemoCommonApproaches()

	printLine()
	fmt.Println()

	// Then show the option pattern solution
	DemoOptionPattern()

	printLine()
	fmt.Println()

	// Finally, show the advanced features
	DemoAdvancedFeatures()

	printLine()
	fmt.Println("  DEMO COMPLETED")
	printLine()
}

// DemoAdvancedFeatures shows how to use functional options to configure a server.
func DemoAdvancedFeatures() {
	fmt.Println("=== Advanced Features ===")

	// Custom logger
	customLogger := log.New(os.Stdout, "CUSTOM: ", log.Lshortfile)

	server, err := NewServer(
		WithHost("production.example.com"),
		WithPort(443),
		WithTLS("/etc/ssl/cert.pem", "/etc/ssl/key.pem"),
		WithMaxConnections(1000),
		WithTimeout(60*time.Second),
		WithLogger(customLogger),
		WithMiddleware("cors", "auth", "ratelimit", "logging"),
		WithReadTimeout(30*time.Second),
		WithWriteTimeout(30*time.Second),
	)

	if err != nil {
		fmt.Printf("✗ Error creating server: %v\n", err)
		return
	}

	fmt.Println("Production server configuration:")
	fmt.Println(server)
	fmt.Println("\nStarting server...")
	server.Start()
}

func printLine() {
	fmt.Println("========================================")
}
