// Package repository provides a generic repository layer for executing database operations
// such as insert, delete, update, get, and get all, using a parser that builds the queries
// and extracts the data.
package repository

import (
	"context"
	"strings"

	"github.com/Class-Connect-GRUPO-5/microservices-common/database"
	"github.com/Class-Connect-GRUPO-5/microservices-common/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository is a generic repository that works with a specific QueryParser to
// execute database operations and return standardized API responses.
type Repository[P QueryParser] struct {
	parser P
	db     *pgxpool.Pool
}

// NewRepository creates a new Repository using the provided QueryParser.
// It connects to the default database configured in the database package.
func NewRepository[P QueryParser](parser P) Repository[P] {
	return Repository[P]{
		parser: parser,
		db:     database.DB,
	}
}

// isBadRequestError determines whether the error is due to invalid input,
// such as violating not-null or check constraints, or SQL syntax issues.
func isBadRequestError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "invalid input") ||
		strings.Contains(msg, "violates not-null constraint") ||
		strings.Contains(msg, "violates check constraint") ||
		strings.Contains(msg, "syntax error")
}

// Insert executes the insert query provided by the parser.
// Returns a 201 Created on success, or an appropriate error response:
// - 409 Conflict if the resource already exists (duplicate key),
// - 400 Bad Request for invalid inputs,
// - 500 Internal Server Error for other failures.
func (r *Repository[P]) Insert() models.APIResponse {
	query, args := r.parser.InsertQuery()
	_, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return models.NewProblemDetails(409, "Conflict", "Resource already exists", "repository.Insert")
		}
		if isBadRequestError(err) {
			return models.NewProblemDetails(400, "Bad Request", err.Error(), "repository.Insert")
		}
		return models.NewProblemDetails(500, "Insert Failed", err.Error(), "repository.Insert")
	}
	return models.NewSuccessDetails(201, "Created", "Insert successful", "repository.Insert", "")
}

// Delete executes the delete query and returns:
// - 200 OK on success,
// - 404 Not Found if no rows were affected,
// - 500 Internal Server Error otherwise.
func (r *Repository[P]) Delete() models.APIResponse {
	query, args := r.parser.DeleteQuery()
	tag, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return models.NewProblemDetails(500, "Delete Failed", err.Error(), "repository.Delete")
	}
	if tag.RowsAffected() == 0 {
		return models.NewProblemDetails(404, "Not Found", "Resource not found", "repository.Delete")
	}
	return models.NewSuccessDetails(200, "Deleted", "Delete successful", "repository.Delete", "")
}

// Update executes the update query and returns:
// - 200 OK if updated successfully,
// - 404 Not Found if no rows were affected,
// - 400 Bad Request if the update failed due to input validation,
// - 500 Internal Server Error otherwise.
func (r *Repository[P]) Update() models.APIResponse {
	query, args := r.parser.UpdateQuery()
	tag, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		if isBadRequestError(err) {
			return models.NewProblemDetails(400, "Bad Request", err.Error(), "repository.Update")
		}
		return models.NewProblemDetails(500, "Update Failed", err.Error(), "repository.Update")
	}
	if tag.RowsAffected() == 0 {
		return models.NewProblemDetails(404, "Not Found", "Resource not found", "repository.Update")
	}
	return models.NewSuccessDetails(200, "Updated", "Update successful", "repository.Update", "")
}

// Get retrieves a single resource using the query from the parser.
// Returns:
// - 200 OK and the result if found,
// - 404 Not Found if no result is returned,
// - 500 Internal Server Error otherwise.
func (r *Repository[P]) Get() models.APIResponse {
	query, args := r.parser.GetQuery()
	row := r.db.QueryRow(context.Background(), query, args...)

	result, err := r.parser.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.NewProblemDetails(404, "Not Found", "Resource not found", "repository.Get")
		}
		return models.NewProblemDetails(500, "Get Failed", err.Error(), "repository.Get")
	}
	return models.NewSuccessDetails(200, "Fetched", "Resource fetched successfully", "repository.Get", result)
}

// GetAll retrieves multiple resources using the parser's query.
// Returns:
// - 200 OK with all results on success,
// - 500 Internal Server Error on failure.
func (r *Repository[P]) GetAll() models.APIResponse {
	query, args := r.parser.GetAllQuery()
	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return models.NewProblemDetails(500, "GetAll Failed", err.Error(), "repository.GetAll")
	}
	defer rows.Close()

	results, err := r.parser.ScanRows(rows)
	if err != nil {
		return models.NewProblemDetails(500, "Scan Failed", err.Error(), "repository.GetAll")
	}
	return models.NewSuccessDetails(200, "Fetched All", "Resources fetched successfully", "repository.GetAll", results)
}
