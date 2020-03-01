package main

import (
	"log"

	"github.com/i-go-go/go-experiment-clickhouse/internal/clickhouse"
	"github.com/i-go-go/go-experiment-clickhouse/internal/server"
)

func main() {
	log.Println("Go Experiment ClickHouse: Demo Server")

	db, err := clickhouse.Init("tcp://clickhouse:9000?debug=true")
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %s", err)
	}

	demoServer := server.InitServer(db)
	if err := demoServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
