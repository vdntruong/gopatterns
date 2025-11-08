package main

import (
	"errors"
	"fmt"
)

type Computer struct {
	CPU     string
	RAM     int
	Storage int
	GPU     string
	OS      string
}

// String returns a string representation of the computer.
// Implement Stringer interface.
func (c *Computer) String() string {
	return fmt.Sprintf("Computer{\n  CPU: %s,\n  RAM: %dGB,\n  Storage: %dGB,\n  GPU: %s,\n  OS: %s\n}",
		c.CPU, c.RAM, c.Storage, c.GPU, c.OS,
	)
}

// ComputerBuilder is a builder for Computer.
type ComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{
		computer: &Computer{},
	}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) SetStorage(storage int) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) SetOS(os string) *ComputerBuilder {
	b.computer.OS = os
	return b
}

// Build validates and returns a Computer.
// Returns an error if validation fails.
func (b *ComputerBuilder) Build() (*Computer, error) {
	// Validate required fields
	if b.computer.CPU == "" {
		return nil, errors.New("CPU is required")
	}
	if b.computer.RAM <= 0 {
		return nil, errors.New("RAM must be greater than 0")
	}
	if b.computer.Storage <= 0 {
		return nil, errors.New("storage must be greater than 0")
	}
	if b.computer.OS == "" {
		return nil, errors.New("OS is required")
	}

	// Optional: Validate reasonable ranges
	if b.computer.RAM > 1024 {
		return nil, errors.New("RAM exceeds maximum allowed (1024GB)")
	}
	if b.computer.Storage > 100000 {
		return nil, errors.New("storage exceeds maximum allowed (100TB)")
	}

	return b.computer, nil
}
