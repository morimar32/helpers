package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb" //mssql implementation
)

// InitConnection initializes a connection to an instance of MSSQL
func InitConnection(constring string, maxOpen int, maxIdle int, maxLifetime time.Duration) (*sql.DB, error) {
	var err error
	con, err := sql.Open("mssql", constring)
	if err != nil {
		return nil, err
	}

	con.SetMaxOpenConns(maxOpen)
	con.SetMaxIdleConns(maxIdle)
	con.SetConnMaxLifetime(maxLifetime)
	err = con.Ping()
	if err != nil {
		return nil, err
	}

	return con, nil
}

// Query Executes the query with the provided query parameters and executes the databind function for each record
func Query(ctx context.Context, con *sql.DB, query string, databind func(rows *sql.Rows) error, queryParams ...interface{}) (int64, error) {
	rows, err := con.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var i int64
	for rows.Next() {
		i++
		if err = databind(rows); err != nil {
			return i, err
		}
	}
	return i, nil
}

// QueryStatement Executes the statement with the provided query parameters and executes the databind function for each record
func QueryStatement(ctx context.Context, stmt *sql.Stmt, databind func(rows *sql.Rows) error, queryParams ...interface{}) (int64, error) {
	rows, err := stmt.QueryContext(ctx, queryParams...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var i int64
	for rows.Next() {
		i++
		if err = databind(rows); err != nil {
			return i, err
		}
	}
	return i, nil
}

// ExecuteNonQuery Executes a command with the given command arguments
func ExecuteNonQuery(ctx context.Context, con *sql.DB, cmd string, cmdArgs ...interface{}) (int64, error) {
	result, err := con.ExecContext(ctx, cmd, cmdArgs...)
	if err != nil {
		return 0, err
	}
	total, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return total, nil
}

// ExecuteStatementNonQuery Executes a command statement with the given command arguments
func ExecuteStatementNonQuery(ctx context.Context, stmt *sql.Stmt, cmdArgs ...interface{}) (int64, error) {
	result, err := stmt.ExecContext(ctx, cmdArgs...)
	if err != nil {
		return 0, err
	}
	total, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return total, nil
}

// NullTimeToTime Converts sql.NullTime to a *time.Time is a value is present, otherwise nil
func NullTimeToTime(val sql.NullTime) *time.Time {
	if !val.Valid {
		return nil
	}
	return &val.Time
}

// GetGUIDString Returns a string representation of a UNIQUEIDENTIFIER from a mssql column, adjusted for endian differences
func GetGUIDString(b []byte) string {
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
