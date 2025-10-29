package main

import (
	"log"
	"time"
)

func NewClientWithLogger(logger *log.Logger) *Client {
	return &Client{
		Logger: logger,
	}
}

func NewClientWithRetries(retries int) *Client {
	return &Client{
		Retries: retries,
	}
}

func NewClientWithTimeout(timeout time.Duration) *Client {
	return &Client{
		Timeout: timeout,
	}
}

func NewClient(logger *log.Logger, retries int, timeout time.Duration) *Client {
	return &Client{
		Logger:  logger,
		Retries: retries,
		Timeout: timeout,
	}
}
