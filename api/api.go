package api

import (
	MIDDLEWARES "test-case-majoo/api/middleware"
	API_V1 "test-case-majoo/api/routers/v1"
	AUTH_REPO "test-case-majoo/infrastructure/repository/auth"
	"test-case-majoo/usecase/auth"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Init(r *gin.Engine, client *gorm.DB) {

	repoAuth := AUTH_REPO.NewAuthRepository(client)
	serviceAuth := auth.NewService(repoAuth)

	router(r, serviceAuth)
}

func router(r *gin.Engine, at *auth.Service) {
	//API V1
	r.Use(MIDDLEWARES.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		API_V1.RouteLogin(v1, at)
	}

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "aku sehat",
		})
	})

}
