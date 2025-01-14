package internal

import "github.com/praveenmahasena/imgserver/internal/server"

func Run() error {
	s := server.New(":42069", "tcp")
	return s.Run()
}
