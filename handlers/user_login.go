package handlers

import (
	"financial_management/model"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	model.User
}

func UserLogin(c *gin.Context) {
	var (
		tmpUser *model.User
		reqBody *UserLoginRequest
	)

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(consts.StatusBadRequest, tmpUser)
		return
	}
}
