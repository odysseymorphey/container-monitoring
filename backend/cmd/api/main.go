package main

import (
	"container-monitoring/internal/api/server"
	_ "github.com/lib/pq"
)

func main() {
	app := server.NewServer()

	app.Run()
}
