package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kungfuxiongmao/cvwo_assign_backend/internal/api"
	"github.com/kungfuxiongmao/cvwo_assign_backend/internal/middleware"
)

func Logout(c *gin.Context) {
	middleware.ClearToken(c)
	api.SuccessMsg(c, nil, "successfully logged out")
}
