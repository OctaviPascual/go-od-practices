package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type ServerOption func(*Server) error

type Server struct {
	timeout time.Duration
}

func NewServer(options ...ServerOption) (*Server, error) {
	s := Server{}
	for _, option := range options {
		if err := option(&s); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	return &s, nil
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) error {
		if timeout == time.Duration(0) {
			return errors.New("timeout cannot be zero")
		}

		s.timeout = timeout
		return nil
	}
}

func (s *Server) Run() {}

func main() {
	server, err := NewServer(WithTimeout(time.Millisecond))
	if err != nil {
		log.Fatalf("could not create new server: %s", err)
	}

	server.Run()
}
