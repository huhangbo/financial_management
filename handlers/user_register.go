package handlers

import (
	"financial_management/consts"
	"financial_management/middleware"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var (
		reqUser *model.User
	)
	if err := c.BindJSON(&reqUser); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if service.IsUserExist(reqUser.Telephone) {
		util.Response(c, consts.UserExistCode, nil)
		return
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
	reqUser.Password = string(hashPassword)
	if err := service.InsertUser(reqUser); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	tmpUser := service.GetUserByTelephone(reqUser.Telephone)
	if tmpUser == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	token := middleware.GenerateToken(tmpUser.UserID)
	tmpUser.Password = ""
	util.Response(c, consts.SuccessCode, UserLoginResp{
		Token: token,
		User:  tmpUser,
	})
}
