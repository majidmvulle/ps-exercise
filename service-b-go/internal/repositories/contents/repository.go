package contents

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/majidmvulle/ps-exercise/service-b-go/internal/models"
)

const defaultPerPage = 20
const maxPerPage = 40

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(ctx context.Context, id uint) (models.ContentWithLocalization, *models.ValidationError) {
	var content models.ContentWithLocalization

	row := r.db.QueryRowContext(ctx, "select * from contents where id=$1", id)
	if err := row.Scan(
		&content.Id,
		&content.Text,
		&content.TextLocalized,
		&content.CreatedAt,
		&content.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ContentWithLocalization{}, &models.ValidationError{
				ErrorCode: models.ErrNotFound,
				Message:   fmt.Errorf("content not found").Error(),
				Errors:    nil,
			}
		}

		return models.ContentWithLocalization{}, &models.ValidationError{
			ErrorCode: models.ValidationErrorErrorCodeInternalError,
			Errors:    nil,
			Message:   err.Error(),
		}
	}

	return content, nil
}

func (r *Repository) List(
	ctx context.Context,
	listParams models.ListContentsParams,
) ([]models.ContentWithLocalization, models.Pagination, *models.ValidationError) {
	contents := make([]models.ContentWithLocalization, 0)

	offset, perPage := offsetLimit(listParams.Page, listParams.PerPage)

	rows, err := r.db.QueryContext(
		ctx,
		"select * from contents offset $1 limit $2", offset, perPage,
	)
	if err != nil {
		return nil,
			models.Pagination{},
			&models.ValidationError{
				ErrorCode: models.ValidationErrorErrorCodeInternalError,
				Errors:    nil,
				Message:   err.Error(),
			}
	}

	defer rows.Close()

	count := 0

	for rows.Next() {
		var content models.ContentWithLocalization

		if err := rows.Scan(
			&content.Id,
			&content.Text,
			&content.TextLocalized,
			&content.CreatedAt,
			&content.UpdatedAt,
		); err != nil {
			return nil,
				models.Pagination{},
				&models.ValidationError{
					ErrorCode: models.ValidationErrorErrorCodeInternalError,
					Errors:    nil,
					Message:   err.Error(),
				}
		}
		contents = append(contents, content)
		count++
	}

	if err := rows.Err(); err != nil {
		return nil,
			models.Pagination{},
			&models.ValidationError{
				ErrorCode: models.ValidationErrorErrorCodeInternalError,
				Errors:    nil,
				Message:   err.Error(),
			}
	}

	total, err := r.total(ctx)
	if err != nil {
		return nil,
			models.Pagination{},
			&models.ValidationError{
				ErrorCode: models.ValidationErrorErrorCodeInternalError,
				Errors:    nil,
				Message:   err.Error(),
			}
	}

	page := 1
	if offset > 0 {
		page = offset/perPage + 1
	}

	totalPages := 1
	if int(total) > perPage {
		totalPages = int(total) / perPage
	}

	return contents,
		models.Pagination{
			Count:      count,
			Page:       page,
			PerPage:    perPage,
			Total:      int(total),
			TotalPages: totalPages,
		},
		nil
}

func (r *Repository) total(ctx context.Context) (int64, error) {
	var total int64

	err := r.db.QueryRowContext(ctx, "select count(*) from contents").Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

func offsetLimit(pg *int, perPg *int) (int, int) {
	offset := 0
	perPage := defaultPerPage

	if perPg != nil {
		perPage = *perPg

		if perPage > maxPerPage {
			perPage = maxPerPage
		}
	}

	if pg != nil && *pg > 0 {
		offset = (*pg - 1) * perPage
	}

	return offset, perPage
}
