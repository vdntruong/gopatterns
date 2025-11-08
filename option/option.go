package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// Server represents an HTTP server configuration
type Server struct {
	host           string
	port           int
	timeout        time.Duration
	maxConnections int
	enableTLS      bool
	certFile       string
	keyFile        string
	logger         *log.Logger
	middleware     []string
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

// ServerOption is a functional option for configuring a Server
type ServerOption func(*Server) error

// NewServer creates a new Server with default values and applies options
func NewServer(opts ...ServerOption) (*Server, error) {
	// Set sensible defaults
	server := &Server{
		host:           "localhost",
		port:           8080,
		timeout:        30 * time.Second,
		maxConnections: 100,
		enableTLS:      false,
		logger:         log.New(os.Stdout, "SERVER: ", log.LstdFlags),
		middleware:     []string{},
		readTimeout:    15 * time.Second,
		writeTimeout:   15 * time.Second,
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Validate configuration
	if err := server.validate(); err != nil {
		return nil, fmt.Errorf("invalid server configuration: %w", err)
	}

	return server, nil
}

// validate checks if the server configuration is valid
func (s *Server) validate() error {
	if s.port < 1 || s.port > 65535 {
		return errors.New("port must be between 1 and 65535")
	}
	if s.maxConnections < 1 {
		return errors.New("maxConnections must be at least 1")
	}
	if s.enableTLS && (s.certFile == "" || s.keyFile == "") {
		return errors.New("TLS enabled but cert/key files not provided")
	}
	return nil
}

// Option functions

// WithHost sets the server host
func WithHost(host string) ServerOption {
	return func(s *Server) error {
		if host == "" {
			return errors.New("host cannot be empty")
		}
		s.host = host
		return nil
	}
}

// WithPort sets the server port
func WithPort(port int) ServerOption {
	return func(s *Server) error {
		if port < 1 || port > 65535 {
			return errors.New("port must be between 1 and 65535")
		}
		s.port = port
		return nil
	}
}

// WithTimeout sets the general timeout
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) error {
		if timeout < 0 {
			return errors.New("timeout cannot be negative")
		}
		s.timeout = timeout
		return nil
	}
}

// WithMaxConnections sets the maximum number of connections
func WithMaxConnections(max int) ServerOption {
	return func(s *Server) error {
		if max < 1 {
			return errors.New("maxConnections must be at least 1")
		}
		s.maxConnections = max
		return nil
	}
}

// WithTLS enables TLS with certificate and key files
func WithTLS(certFile, keyFile string) ServerOption {
	return func(s *Server) error {
		if certFile == "" || keyFile == "" {
			return errors.New("cert and key files must be provided")
		}
		s.enableTLS = true
		s.certFile = certFile
		s.keyFile = keyFile
		return nil
	}
}

// WithLogger sets a custom logger
func WithLogger(logger *log.Logger) ServerOption {
	return func(s *Server) error {
		if logger == nil {
			return errors.New("logger cannot be nil")
		}
		s.logger = logger
		return nil
	}
}

// WithMiddleware adds middleware to the server
func WithMiddleware(middleware ...string) ServerOption {
	return func(s *Server) error {
		s.middleware = append(s.middleware, middleware...)
		return nil
	}
}

// WithReadTimeout sets the read timeout
func WithReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) error {
		if timeout < 0 {
			return errors.New("read timeout cannot be negative")
		}
		s.readTimeout = timeout
		return nil
	}
}

// WithWriteTimeout sets the write timeout
func WithWriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) error {
		if timeout < 0 {
			return errors.New("write timeout cannot be negative")
		}
		s.writeTimeout = timeout
		return nil
	}
}

// Getter methods (since fields are private)

func (s *Server) Host() string                { return s.host }
func (s *Server) Port() int                   { return s.port }
func (s *Server) Timeout() time.Duration      { return s.timeout }
func (s *Server) MaxConnections() int         { return s.maxConnections }
func (s *Server) EnableTLS() bool             { return s.enableTLS }
func (s *Server) CertFile() string            { return s.certFile }
func (s *Server) KeyFile() string             { return s.keyFile }
func (s *Server) Logger() *log.Logger         { return s.logger }
func (s *Server) Middleware() []string        { return s.middleware }
func (s *Server) ReadTimeout() time.Duration  { return s.readTimeout }
func (s *Server) WriteTimeout() time.Duration { return s.writeTimeout }

func (s *Server) String() string {
	tlsStatus := "disabled"
	if s.enableTLS {
		tlsStatus = "enabled"
	}
	return fmt.Sprintf("Server{Host: %s, Port: %d, TLS: %s, MaxConns: %d, Timeout: %v, Middleware: %v}",
		s.host, s.port, tlsStatus, s.maxConnections, s.timeout, s.middleware)
}

// Start simulates starting the server
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s:%d", s.host, s.port)
	s.logger.Printf("TLS: %v, Max Connections: %d", s.enableTLS, s.maxConnections)
	s.logger.Printf("Timeouts - General: %v, Read: %v, Write: %v",
		s.timeout, s.readTimeout, s.writeTimeout)
	if len(s.middleware) > 0 {
		s.logger.Printf("Middleware: %v", s.middleware)
	}
	return nil
}

// Advanced example: Composable options

// WithProduction is a composite option that configures a production-ready server
func WithProduction() ServerOption {
	return func(s *Server) error {
		// Apply multiple configurations
		opts := []ServerOption{
			WithHost("0.0.0.0"),
			WithMaxConnections(1000),
			WithTimeout(60 * time.Second),
			WithReadTimeout(30 * time.Second),
			WithWriteTimeout(30 * time.Second),
			WithMiddleware("logging", "recovery", "compression"),
		}

		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithDevelopment is a composite option for development environment
func WithDevelopment() ServerOption {
	return func(s *Server) error {
		opts := []ServerOption{
			WithHost("localhost"),
			WithPort(3000),
			WithMaxConnections(10),
			WithTimeout(5 * time.Second),
			WithMiddleware("logging", "debug"),
		}

		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}
		return nil
	}
}

// DemoOptionPattern shows how to use functional options to configure a server.
func DemoOptionPattern() {
	fmt.Println("=== Option Pattern Examples ===")
	fmt.Println()

	// Example 1: Simple server with defaults
	server1, err := NewServer()
	if err != nil {
		panic(err)
	}
	fmt.Println("1. Server with defaults:")
	fmt.Println(server1)
	fmt.Println()

	// Example 2: Custom configuration
	server2, err := NewServer(
		WithHost("api.example.com"),
		WithPort(443),
		WithTLS("/path/to/cert.pem", "/path/to/key.pem"),
		WithMaxConnections(500),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("2. Custom HTTPS server:")
	fmt.Println(server2)
	fmt.Println()

	// Example 3: With middleware
	server3, err := NewServer(
		WithPort(8000),
		WithMiddleware("cors", "auth", "ratelimit"),
		WithTimeout(45*time.Second),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("3. Server with middleware:")
	fmt.Println(server3)
	fmt.Println()

	// Example 4: Production preset
	server4, err := NewServer(
		WithProduction(),
		WithPort(443),
		WithTLS("/path/to/cert.pem", "/path/to/key.pem"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("4. Production server:")
	fmt.Println(server4)
	server4.Start()
	fmt.Println()

	// Example 5: Development preset
	server5, err := NewServer(WithDevelopment())
	if err != nil {
		panic(err)
	}
	fmt.Println("5. Development server:")
	fmt.Println(server5)
	fmt.Println()

	// Example 6: Error handling - invalid configuration
	fmt.Println("6. Error handling example:")
	_, err = NewServer(
		WithPort(99999), // Invalid port
	)
	if err != nil {
		fmt.Printf("✓ Validation caught error: %v\n", err)
	}
	fmt.Println()

	// Example 7: TLS without certificates
	_, err = NewServer(
		WithTLS("", ""), // Missing cert files
	)
	if err != nil {
		fmt.Printf("✓ Validation caught error: %v\n", err)
	}
}
