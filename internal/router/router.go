package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/cvwo_assign_backend/internal/middleware"
	"github.com/kungfuxiongmao/cvwo_assign_backend/internal/routes"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsSetUp())
	r.Use(middleware.DBToContext(db))
	routes.GetRoutes(r)
	return r
}
