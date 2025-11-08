package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Database represents a database connection configuration
type Database struct {
	Host           string
	Port           int
	Username       string
	Password       string
	MaxConnections int
	ConnTimeout    time.Duration
	IdleTimeout    time.Duration
	EnableSSL      bool
	Logger         *log.Logger
	RetryAttempts  int
	ConnectionPool bool
}

// NewDatabase is a factory method for creating a new Database instance
//
// Approach 1: Simple constructor with all parameters
// Problem: Too many parameters, hard to remember order, error-prone - SornaQube failed for this
func NewDatabase(host string, port int, username, password string, maxConns int,
	connTimeout, idleTimeout time.Duration, enableSSL bool, logger *log.Logger,
	retryAttempts int, connectionPool bool) *Database {
	return &Database{
		Host:           host,
		Port:           port,
		Username:       username,
		Password:       password,
		MaxConnections: maxConns,
		ConnTimeout:    connTimeout,
		IdleTimeout:    idleTimeout,
		EnableSSL:      enableSSL,
		Logger:         logger,
		RetryAttempts:  retryAttempts,
		ConnectionPool: connectionPool,
	}
}

// NewDatabaseWithDefaults is a factory method for creating a new Database instance with default values
//
// Approach 2: Constructor with default values
// Problem: Inflexible - can't override specific defaults without new constructors
func NewDatabaseWithDefaults(host string, port int, username, password string) *Database {
	return &Database{
		Host:           host,
		Port:           port,
		Username:       username,
		Password:       password,
		MaxConnections: 100,
		ConnTimeout:    30 * time.Second,
		IdleTimeout:    5 * time.Minute,
		EnableSSL:      true,
		Logger:         log.Default(),
		RetryAttempts:  3,
		ConnectionPool: true,
	}
}

// NewDatabaseBasic is a factory method for creating a new Database instance with basic configuration
//
// Approach 3: Telescoping constructors
// Problem: Exponential growth - need many constructors for different combinations
func NewDatabaseBasic(host string, port int) *Database {
	return &Database{
		Host: host,
		Port: port,
	}
}

func NewDatabaseWithAuth(host string, port int, username, password string) *Database {
	return &Database{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func NewDatabaseWithSSL(host string, port int, username, password string, enableSSL bool) *Database {
	return &Database{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		EnableSSL: enableSSL,
	}
}

func NewDatabaseWithPooling(host string, port int, username, password string,
	enableSSL bool, maxConns int, connectionPool bool) *Database {
	return &Database{
		Host:           host,
		Port:           port,
		Username:       username,
		Password:       password,
		EnableSSL:      enableSSL,
		MaxConnections: maxConns,
		ConnectionPool: connectionPool,
	}
}

// DatabaseConfig represents a database configuration
//
// Approach 4: Config struct
// Problem: Still verbose, no validation, easy to forget required fields
type DatabaseConfig struct {
	Host           string
	Port           int
	Username       string
	Password       string
	MaxConnections int
	ConnTimeout    time.Duration
	IdleTimeout    time.Duration
	EnableSSL      bool
	Logger         *log.Logger
	RetryAttempts  int
	ConnectionPool bool
}

func NewDatabaseFromConfig(config DatabaseConfig) *Database {
	return &Database{
		Host:           config.Host,
		Port:           config.Port,
		Username:       config.Username,
		Password:       config.Password,
		MaxConnections: config.MaxConnections,
		ConnTimeout:    config.ConnTimeout,
		IdleTimeout:    config.IdleTimeout,
		EnableSSL:      config.EnableSSL,
		Logger:         config.Logger,
		RetryAttempts:  config.RetryAttempts,
		ConnectionPool: config.ConnectionPool,
	}
}

// NewEmptyDatabase is a factory method for creating a new Database instance with default values
//
// Approach 5: Setters after construction
// Problem: Object mutable, no immutability, can be used in invalid state
func NewEmptyDatabase() *Database {
	return &Database{}
}

func (d *Database) SetHost(host string) {
	d.Host = host
}

func (d *Database) SetPort(port int) {
	d.Port = port
}

func (d *Database) SetAuth(username, password string) {
	d.Username = username
	d.Password = password
}

func (d *Database) SetMaxConnections(max int) {
	d.MaxConnections = max
}

func (d *Database) String() string {
	return fmt.Sprintf("Database{Host: %s, Port: %d, User: %s, MaxConns: %d, SSL: %v, Pool: %v, Timeout: %v}",
		d.Host, d.Port, d.Username, d.MaxConnections, d.EnableSSL, d.ConnectionPool, d.ConnTimeout)
}

// DemoCommonApproaches demonstrates the problems with traditional approaches to creating objects
//
// Example usage demonstrating the problems with traditional approaches
func DemoCommonApproaches() {
	fmt.Println("=== Common Constructor Approaches (Without Option Pattern) ===")

	// Approach 1: All parameters - extremely hard to read and error-prone
	db1 := NewDatabase(
		"localhost", 5432, "admin", "secret123", 100,
		30*time.Second, 5*time.Minute, true, log.Default(), 3, true,
	)
	fmt.Println("1. All parameters constructor:")
	fmt.Println(db1)
	fmt.Println("Problem: Which parameter is which? Easy to pass wrong values!")

	// Approach 2: With defaults - can't customize easily
	db2 := NewDatabaseWithDefaults("postgres.example.com", 5432, "user", "pass")
	fmt.Println("2. Constructor with defaults:")
	fmt.Println(db2)
	fmt.Println("Problem: What if I want SSL disabled but keep other defaults?")

	// Approach 3: Telescoping - need many constructors
	db3 := NewDatabaseWithSSL("db.local", 3306, "root", "password", false)
	fmt.Println("3. Telescoping constructor:")
	fmt.Println(db3)
	fmt.Println("Problem: Need separate constructor for every combination!")

	// Approach 4: Config struct - verbose
	config := DatabaseConfig{
		Host:           "cache.example.com",
		Port:           6379,
		Username:       "redis",
		Password:       "redis123",
		MaxConnections: 50,
		ConnTimeout:    10 * time.Second,
		EnableSSL:      false,
		Logger:         log.New(os.Stdout, "DB: ", log.LstdFlags),
		ConnectionPool: true,
	}
	db4 := NewDatabaseFromConfig(config)
	fmt.Println("4. Config struct approach:")
	fmt.Println(db4)
	fmt.Println("Problem: Verbose, no validation, easy to forget fields")

	// Approach 5: Setters - object can be in invalid state
	db5 := NewEmptyDatabase()
	db5.SetHost("mysql.local")
	db5.SetPort(3306)
	// Oops, forgot to set auth! Database is in invalid state
	fmt.Println("5. Setters approach:")
	fmt.Println(db5)
	fmt.Println("Problem: Object can be used before fully configured!")

	fmt.Println("=== The Option Pattern solves all these problems! ===")
}
