package main

import (
	"io"
	"log"
	"time"

	"github.com/i-go-go/go-experiment-clickhouse/internal/clickhouse"
	"github.com/i-go-go/go-experiment-clickhouse/internal/csv"
	"github.com/i-go-go/go-experiment-clickhouse/internal/entity"
	"github.com/jmoiron/sqlx"
)

func main() {
	log.Println("Go Experiment ClickHouse: Populate database")

	// "tcp://clickhouse:9000?debug=true"
	db := connectToClickHouse("tcp://localhost:9000?debug=true")
	populateHotelBookings("data/hotel_bookings_small.csv", db)
}

func connectToClickHouse(dsn string) *sqlx.DB {
	db, err := clickhouse.Init(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %s", err)
	}

	log.Print("Successful connected to ClickHouse")

	// Prepare database
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS hotels"); err != nil {
		log.Fatal("Failed to create database", err)
	}

	// ToDo: Prepare table

	return db
}

func populateHotelBookings(csvFile string, _ *sqlx.DB) {
	parser, err := csv.Parse(csvFile, true)
	if err != nil {
		log.Fatalf("Failed to load CSV file: %s", err)
	}

	for {
		row, err := parser.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Failed to read data from CSV", err)
		}

		hotelBooking := &entity.HotelBooking{}
		if err := parser.FillStruct(row, hotelBooking); err != nil {
			log.Fatalf("Failed to fill struct: %s", err)
		}

		log.Printf("HotelBooking: %+v", row)
		time.Sleep(time.Second)

		// ToDo: Populate data to ClickHouse (use db)
	}
}
