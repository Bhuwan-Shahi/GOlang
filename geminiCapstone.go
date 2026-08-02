package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrServerTimeout     = errors.New("server health check timed out")
	ErrServerUnreachable = errors.New("server is offline or unreachable")
)

type server struct {
	ID          int
	Host        string
	ActiveConns int
}
type HealthResult struct {
	ServerID int
	Success  bool
	Conns    int
	Err      error
}
type Dashboard struct {
	mu               sync.Mutex
	totalConns       int
	healthyServers   int
	unhealthyServers int
}

func (d *Dashboard) RecordResult(success bool, conns int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if success {
		d.healthyServers++
		d.totalConns += conns
	} else {
		d.unhealthyServers++
	}
}
func main() {
	server1 := server{2, "Bhuwan", 3000}
	fmt.Println(server1)
}
