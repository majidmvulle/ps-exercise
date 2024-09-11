package contents

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majidmvulle/ps-exercise/service-b-go/internal/models"
)

type contentsRepo interface {
	Get(ctx context.Context, id uint) (models.ContentWithLocalization, *models.ValidationError)
	List(
		ctx context.Context,
		params models.ListContentsParams,
	) ([]models.ContentWithLocalization, models.Pagination, *models.ValidationError)
}

type Controller struct {
	repo contentsRepo
}

func New(repo contentsRepo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (ct *Controller) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)

		return
	}

	content, validationErr := ct.repo.Get(c.Request.Context(), uint(id))
	if validationErr != nil {
		status := http.StatusUnprocessableEntity

		if validationErr.ErrorCode == models.ErrNotFound {
			status = http.StatusNotFound
		}

		c.JSON(status, validationErr)

		return
	}

	c.JSON(http.StatusOK, content)
}

func (ct *Controller) List(c *gin.Context) {
	listParams := models.ListContentsParams{}
	if err := c.ShouldBindQuery(&listParams); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)

		return
	}

	res, pagination, err := ct.repo.List(c, listParams)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)

		return
	}

	contents := make([]models.Content, len(res), len(res))

	for i, content := range res {
		contents[i] = models.Content{
			Text:      content.Content.Text,
			CreatedAt: content.Content.CreatedAt,
			Id:        content.Content.Id,
			UpdatedAt: content.Content.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, models.ListContents{
		Data: contents,
		Metadata: models.ListMetadata{
			Pagination: pagination,
		},
	})
}
