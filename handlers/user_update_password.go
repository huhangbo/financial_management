package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordReq struct {
	UserID      int     `json:"user_id"`
	OldPassword *string `json:"old_password"`
	NewPassword string  `json:"new_password"`
}

func UpdatePassword(c *gin.Context) {
	var (
		req *UpdatePasswordReq
		uid = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&req); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpUser := service.GetUserByID(req.UserID)

	if uid != setting.Config.Admin.ID {
		if uid != req.UserID {
			util.Response(c, consts.PermissionErrorCode, nil)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(tmpUser.Password), []byte(*req.OldPassword)); err != nil {
			util.Response(c, consts.PasswordErrorCode, nil)
			return
		}
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	tmpUser.Password = string(hashPassword)
	if err := service.UpdateUser(tmpUser); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
