package option

// 方便創建一個帶有default值的object, 並可以選擇性的修改一些object的參數

import (
	"time"
)

// 創建一個帶有default值的server object, 並可選擇timeout,log level等...設定

type Server struct {
	addr     string
	port     int
	timeout  time.Duration
	logLevel string
}

type Option func(*Server)

func WithTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.timeout = t
	}
}

func WithLogLevel(level string) Option {
	return func(s *Server) {
		s.logLevel = level
	}
}

func NewServer(addr string, port int, opts ...Option) *Server {
	s := &Server{
		addr:     addr,
		port:     port,
		timeout:  10 * time.Second,
		logLevel: "info",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
