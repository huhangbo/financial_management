package handlers

import (
	"financial_management/consts"
	"financial_management/middleware"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserLoginResp struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func UserLogin(c *gin.Context) {
	var (
		reqUser *model.User
	)

	if err := c.BindJSON(&reqUser); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if !service.IsUserExist(reqUser.Telephone) {
		util.Response(c, consts.UserNotExistCode, nil)
		return
	}

	tmpUser := service.GetUserByTelephone(reqUser.Telephone)
	if tmpUser == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tmpUser.Password), []byte(reqUser.Password)); err != nil {
		util.Response(c, consts.PasswordErrorCode, nil)
		return
	}
	tmpUser.Password = ""
	token := middleware.GenerateToken(tmpUser.UserID)
	util.Response(c, consts.SuccessCode, UserLoginResp{
		Token: token,
		User:  tmpUser,
	})
}
