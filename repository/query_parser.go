package repository

import "github.com/jackc/pgx/v5"

// QueryParser defines an interface for building SQL queries from custom input.
// Each method should return the SQL query string and the arguments to bind.
type QueryParser interface {
	InsertQuery(data any) (string, []any)
	UpdateQuery(data any) (string, []any)
	DeleteQuery(id string) (string, []any)
	GetQuery(id string) (string, []any)
	GetAllQuery() (string, []any)
	ScanRow(row pgx.Row) (string, error)
	ScanRows(rows pgx.Rows) (string, error)
}
