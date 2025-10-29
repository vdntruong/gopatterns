package main

import (
	"log"
	"time"
)

type Client struct {
	Logger  *log.Logger
	Timeout time.Duration
	Retries int
}

func main() {
	// Create a client with default settings
	defaultClient := NewClientWithOptions()
	log.Printf("Default Client: Timeout=%v, Retries=%d", defaultClient.Timeout, defaultClient.Retries)

	defaultClient.Apply(
		WithLogger(log.New(log.Writer(), "MyApp: ", log.LstdFlags)),
	)
	defaultClient.Logger.Println("Applied logger")

	// Create a client with custom timeout and retries
	customClient := NewClientWithOptions(
		WithTimeout(30*time.Second),
		WithRetries(5),
	)
	log.Printf("Custom Client: Timeout=%v, Retries=%d", customClient.Timeout, customClient.Retries)
}
