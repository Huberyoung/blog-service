package middleware

import (
	"blog-service/global"
	"blog-service/pkg/app"
	"blog-service/pkg/email"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	failMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.CloneWithCallersFrames().ErrorF("panic recover err:%v", err)

				err := failMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err))
				if err != nil {
					global.Logger.PanicF("mail.SendMail err:%v", err)
				}
				app.NewResponse(c).ToErrorResponseList(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
