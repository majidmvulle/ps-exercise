package health

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/majidmvulle/ps-exercise/service-b-go/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Ping(ctx context.Context) models.HealthCheck {
	err := r.db.PingContext(ctx)
	if err != nil {
		return models.HealthCheck{
			Status:  models.DOWN,
			Message: err.Error(),
		}
	}

	return models.HealthCheck{
		Message: fmt.Sprintf("PONG"),
		Status:  models.UP,
	}
}
