package clickhouse

import (
	"github.com/jmoiron/sqlx"

	// ClickHouse driver
	_ "github.com/ClickHouse/clickhouse-go"
)

func Init(dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
