package main

import (
	"time"
)

type ServerOption func(*Server)

type Server struct {
	timeout time.Duration
}

func NewServer(options ...ServerOption) *Server {
	s := Server{}
	for _, option := range options {
		option(&s)
	}

	return &s
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func (s *Server) Run() {}

func main() {
	server := NewServer(WithTimeout(time.Millisecond))

	server.Run()
}
