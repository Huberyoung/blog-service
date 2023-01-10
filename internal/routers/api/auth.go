package api

import (
	"github.com/gin-gonic/gin"
	"service/global"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/errorcode"
)

func GetAuth(ctx *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(ctx.Request.Context())
	err := srv.CheckAuth(&param)
	if err != nil {
		global.Logger.ErrorF("srv.CheckAuth errs:%v", errs)
		response.ToErrorResponse(errorcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.ErrorF("srv.GenerateToken errs:%v", errs)
		response.ToErrorResponse(errorcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
