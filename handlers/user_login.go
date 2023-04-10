package handlers

import (
	"context"
	"encoding/json"
	"financial_management/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserLoginRequest struct {
	model.User
}

func UserLogin(context context.Context, c *app.RequestContext) {
	var (
		tmpUser *model.User
		reqBody *UserLoginRequest
	)
	body, _ := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		c.JSON(consts.StatusBadRequest, tmpUser)
		return
	}
}
