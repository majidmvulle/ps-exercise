package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	ctrlContents "github.com/majidmvulle/ps-exercise/service-b-go/internal/controllers/contents"
	ctrlHealth "github.com/majidmvulle/ps-exercise/service-b-go/internal/controllers/health"
	repoContents "github.com/majidmvulle/ps-exercise/service-b-go/internal/repositories/contents"
	repoHealth "github.com/majidmvulle/ps-exercise/service-b-go/internal/repositories/health"
)

func contentsRoutes(db *sql.DB, group *gin.RouterGroup) {
	ctrl := ctrlContents.New(repoContents.New(db))

	grp := group.Group("/contents")
	{
		grp.GET("", ctrl.List)

		id := grp.Group("/:id")
		{
			id.GET("", ctrl.Get)
		}
	}
}

func healthRoutes(db *sql.DB, group *gin.Engine) {
	ctrl := ctrlHealth.New(repoHealth.New(db))

	group.GET("/health", ctrl.Get)
}
