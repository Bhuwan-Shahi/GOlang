package main

import (
	"errors"
	"fmt"
)

var ServerUnreachable = errors.New("The server is unreachange!")

type Server struct {
	id     int
	count  int
	status string
}

func (server *Server) trackStatus() error {
	if server == nil {
		return errors.New("Nile pointer server")
	}
	if server.status == "offline" {
		return errors.New("Server is currenlty offline")
	}
	if server.status == "unavailable" {
		return fmt.Errorf("The server with id %s: %e", server.id, ServerUnreachable)
	}
	fmt.Println("Server of id ", server.id, "is ", server.status, "with", server.count, "connection")
	return nil
}
func worker(server *Server, ch chan error) {
	ch <- server.trackStatus()
}
func main() {
	ch := make(chan error)
	s1 := Server{}
	s2 := Server{1, 23, "online"}
	s3 := Server{2, 0, "unavailabe"}

	go worker(&s1, ch)
	go worker(&s2, ch)
	go worker(&s3, ch)

	err := <-ch
	err2 := <-ch
	err1 := <-ch
	if errors.Is(ServerUnreachable, err1) {
		fmt.Println("Caught:", err1)
	}
	if errors.Is(ServerUnreachable, err) {
		fmt.Println("Caught:", err)
	}
	if errors.Is(ServerUnreachable, err2) {
		fmt.Println("Caught:", err2)
	}

}
