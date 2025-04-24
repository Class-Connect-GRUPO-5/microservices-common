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

type Repository[P QueryParser] struct {
	parser P
	db     *pgxpool.Pool
}

func NewRepository[P QueryParser](parser P) Repository[P] {
	return Repository[P]{
		parser: parser,
		db:     database.DB,
	}
}

func isBadRequestError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "invalid input") ||
		strings.Contains(msg, "violates not-null constraint") ||
		strings.Contains(msg, "violates check constraint") ||
		strings.Contains(msg, "syntax error")
}

func (r *Repository[P]) Insert(data any) models.APIResponse {
	query, args := r.parser.InsertQuery(data)
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

func (r *Repository[P]) Update(data any) models.APIResponse {
	query, args := r.parser.UpdateQuery(data)
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

func (r *Repository[P]) Delete(id string) models.APIResponse {
	query, args := r.parser.DeleteQuery(id)
	tag, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return models.NewProblemDetails(500, "Delete Failed", err.Error(), "repository.Delete")
	}
	if tag.RowsAffected() == 0 {
		return models.NewProblemDetails(404, "Not Found", "Resource not found", "repository.Delete")
	}
	return models.NewSuccessDetails(200, "Deleted", "Delete successful", "repository.Delete", "")
}

func (r *Repository[P]) GetByMany(filters map[string]any) models.APIResponse {
	query, args := r.parser.GetQueryMany(filters)
	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return models.NewProblemDetails(500, "GetByMany Failed", err.Error(), "repository.GetByMany")
	}
	defer rows.Close()

	results, err := r.parser.ScanRows(rows)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.NewProblemDetails(404, "Not Found", "Resource not found", "repository.Get")
		}
		return models.NewProblemDetails(500, "Scan Failed", err.Error(), "repository.GetByMany")
	}
	return models.NewSuccessDetails(200, "Fetched", "Resources fetched successfully", "repository.GetByMany", results)
}

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
