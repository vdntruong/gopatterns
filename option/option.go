package main

import (
	"log"
	"time"
)

func NewClientWithOptions(options ...ClientOption) *Client {
	client := &Client{
		Timeout: 10 * time.Second, // Default timeout
		Logger:  log.Default(),    // Default logger
		Retries: 3,                // Default retries
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

func (c *Client) Apply(options ...ClientOption) {
	for _, opt := range options {
		opt(c)
	}
}

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

func WithLogger(logger *log.Logger) ClientOption {
	return func(c *Client) {
		c.Logger = logger
	}
}

func WithRetries(retries int) ClientOption {
	return func(c *Client) {
		c.Retries = retries
	}
}
