package storage

import (
	"container-monitoring/internal/api/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	storage *sql.DB
}

func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	const op = "storage.NewConnection"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Printf("connStr: %s\n", connStr)

	var err error
	storage, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open database: %v", op, err)
	}

	if err := storage.Ping(); err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %v", op, err)
	}

	fmt.Printf("Open database: %s\n", op)

	return &PostgresDB{
		storage: storage,
	}, nil
}

func (db *PostgresDB) GetStatuses(ctx context.Context) ([]models.PingStatus, error) {
	const op = "storage.GetStatuses"

	query := `SELECT ip_address, ping_time, success, last_successful_ping FROM ping_status`
	rows, err := db.storage.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("%s: failed to query database: %v", op, err)
	}
	defer rows.Close()

	var statuses []models.PingStatus
	for rows.Next() {
		var s models.PingStatus
		if err := rows.Scan(&s.IPAddress, &s.PingTime, &s.Success, &s.LastSuccess); err != nil {
			fmt.Printf("%s: failed to scan row: %v", op, err)
		}
		statuses = append(statuses, s)
	}

	return statuses, nil
}

func (db *PostgresDB) AddStatus(ctx context.Context, status *models.PingStatus) error {
	const op = "storage.AddStatus"

	query := `INSERT INTO ping_status (ip_address, ping_time, success, last_successful_ping)
			  VALUES ($1, $2, $3, $4)`
	_, err := db.storage.ExecContext(ctx, query, status.IPAddress, status.PingTime, status.Success, status.LastSuccess)
	if err != nil {
		return fmt.Errorf("%s: failed to insert ping status: %v", op, err)
	}

	return nil
}

func (db *PostgresDB) Ping() error {
	const op = "storage.Ping"

	if err := db.storage.Ping(); err != nil {
		return fmt.Errorf("%s: failed to ping database: %v", op, err)
	}

	return nil
}

func (db *PostgresDB) Close() error {
	const op = "storage.Close"

	if err := db.storage.Close(); err != nil {
		return fmt.Errorf("%s: failed to close database: %v", op, err)
	}

	return nil
}
