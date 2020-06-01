package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" //mssql implementation
)

// DbHelper Helper class for working with databases
type DbHelper struct {
	ConnectionString string
}

// Query Executes the query with the provided query parameters and executes the databind function for each record
func (db *DbHelper) Query(ctx context.Context, query string, databind func(rows *sql.Rows) error, queryParams ...interface{}) error {
	con, err := sql.Open("mssql", db.ConnectionString)
	if err != nil {
		return err
	}
	defer con.Close()

	rows, err := con.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err = databind(rows); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteNonQuery Executes a command with the given command arguments
func (db *DbHelper) ExecuteNonQuery(ctx context.Context, cmd string, cmdArgs ...interface{}) (int64, error) {
	con, err := sql.Open("mssql", db.ConnectionString)
	if err != nil {
		return 0, err
	}
	defer con.Close()

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

// GetGUIDString Returns a string representation of a UNIQUEIDENTIFIER from a mssql column, adjusted for endian differences
func GetGUIDString(b []byte) string {
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
