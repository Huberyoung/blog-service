package middleware

import (
	"blog-service/global"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/universal-translator"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	zhtwtranslations "github.com/go-playground/validator/v10/translations/zh_tw"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		var err error
		if ok {
			switch locale {
			case "en":
				err = entranslations.RegisterDefaultTranslations(v, trans)
			case "zh":
				err = zhtranslations.RegisterDefaultTranslations(v, trans)
			case "zh_Hant_TW":
				err = zhtwtranslations.RegisterDefaultTranslations(v, trans)
			default:
				err = zhtranslations.RegisterDefaultTranslations(v, trans)
			}
			if err != nil {
				global.Logger.WarningF("语言包转换错误:%s\n", err.Error())
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
