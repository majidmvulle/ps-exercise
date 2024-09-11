package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/majidmvulle/ps-exercise/service-b-go/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.Config()
	ctx := context.Background()

	db, err := sql.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(cfg.Gin.Mode)

	router := gin.Default()

	router.Use(cors.Default())
	router.Use(AuthApiKey(cfg.ApiKey))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			contentsRoutes(db, v1)
		}
	}

	healthRoutes(db, router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		default:
			return router.Run(fmt.Sprintf(":%d", cfg.Gin.Port))
		}
	})

	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server terminated")
}

func AuthApiKey(cfgApiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-API-Key")

		if apiKey != cfgApiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error_code": "unauthorized"})

			return
		}

		c.Next()
	}
}
