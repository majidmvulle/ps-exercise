package health

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/majidmvulle/ps-exercise/service-b-go/internal/models"
	"net/http"
)

type healthRepo interface {
	Ping(ctx context.Context) models.HealthCheck
}

type Controller struct {
	repo healthRepo
}

func New(repo healthRepo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (ct *Controller) Get(c *gin.Context) {
	status := http.StatusOK

	healthCheck := ct.repo.Ping(c.Request.Context())
	if healthCheck.Status != models.UP {
		status = http.StatusBadRequest
	}

	c.JSON(status, healthCheck)
}
