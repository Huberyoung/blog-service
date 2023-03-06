package v1

import (
	"blog-service/global"
	"blog-service/interbal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// Get godoc
//
//	@Summary		获取账号的token
//	@Description	通过账号名称和密码获得token
//	@Tags			User
//	@Produce		json
//	@Param			username	formData	string			true	"账号昵称"	minlength(1)	maxlength(10)
//	@Param			password	formData	string			true	"账号密码"	minlength(3)	maxlength(100)
//	@Success		200			{object}	model.GetUser	"成功"
//	@Failure		400			{object}	errcode.Error	"请求错误"
//	@Failure		404			{object}	errcode.Error	"找不到页面"
//	@Failure		500			{object}	errcode.Error	"内部错误"
//	@Router			/user [post]
func (u User) Get(c *gin.Context) {
	param := service.GetUserRequest{}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errors)
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckUser(&param)
	if err != nil {
		global.Logger.ErrorF("svc.CheckUser err:%v", err)
		response.ToErrorResponseList(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Username, param.Password)
	if err != nil {
		global.Logger.ErrorF("app.GenerateToken err:%v", err)
		response.ToErrorResponseList(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
