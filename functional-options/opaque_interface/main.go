package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const defaultTimeout = time.Minute

type ServerOption interface {
	apply(*serverOptions) error
}

type serverOption func(*serverOptions) error

func (o serverOption) apply(s *serverOptions) error {
	return o(s)
}

type serverOptions struct {
	localMode bool
	timeout   time.Duration
}

type Server struct {
	host    string
	timeout time.Duration
}

func NewServer(options ...ServerOption) (*Server, error) {
	serverOptions := serverOptions{
		timeout: defaultTimeout,
	}
	if err := WithServerOptions(options...).apply(&serverOptions); err != nil {
		return nil, err
	}

	host := "host"
	if serverOptions.localMode {
		host = "localhost"
	}

	return &Server{
		host:    host,
		timeout: serverOptions.timeout,
	}, nil
}

func WithTimeout(timeout time.Duration) ServerOption {
	f := func(s *serverOptions) error {
		if timeout == time.Duration(0) {
			return errors.New("timeout cannot be zero")
		}

		s.timeout = timeout
		return nil
	}
	return serverOption(f)
}

func WithLocalMode() ServerOption {
	f := func(s *serverOptions) error {
		s.localMode = true
		return nil
	}
	return serverOption(f)
}

func WithServerOptions(options ...ServerOption) ServerOption {
	f := func(s *serverOptions) error {
		for _, option := range options {
			if err := option.apply(s); err != nil {
				return fmt.Errorf("error applying option: %w", err)
			}
		}
		return nil
	}
	return serverOption(f)
}

func (s *Server) Run() {}

func main() {
	server, err := NewServer(WithTimeout(time.Millisecond), WithLocalMode())
	if err != nil {
		log.Fatalf("could not create new server: %s", err)
	}

	server.Run()
}
