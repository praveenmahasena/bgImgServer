package server

import (
	"fmt"
	"log"
	"net"
	"os"
)

type port string

type Server struct {
	Addr port
	Con  string
}

func New(p port, con string) Server {
	return Server{
		p,
		con,
	}
}

func (s Server) Run() error {
	li, err := net.Listen(s.Con, string(s.Addr))
	if err != nil {
		return fmt.Errorf("error during starting the server %v", err)
	}
	f, fErr := os.Create("./log")

	if fErr != nil {
		return fmt.Errorf("error during opening log file")
	}

	defer f.Close()

	l := log.Default()
	l.SetOutput(f)

	for {
		con, conErr := li.Accept()

		if conErr != nil {
			// do better error handling with slog
			continue
		}

		go handle(con, l)
	}
}

func handle(c net.Conn, l *log.Logger) {
	defer c.Close()
	b, err := getImg()
	if err != nil {
		c.Write(nil)
		l.Println(err)
		return
	}
	c.Write(b)
}
