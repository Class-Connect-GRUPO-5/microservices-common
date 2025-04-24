package repository

import "github.com/jackc/pgx/v5"

// QueryParser defines an interface for building SQL queries from custom input.
// Each method should return the SQL query string and the arguments to bind.
type QueryParser interface {
	InsertQuery() (string, []any)
	UpdateQuery() (string, []any)
	DeleteQuery() (string, []any)
	GetQuery() (string, []any)
	GetAllQuery() (string, []any)
	ScanRow(row pgx.Row) (string, error)
	ScanRows(rows pgx.Rows) (string, error)
}
