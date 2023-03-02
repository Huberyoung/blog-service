package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var errs ValidErrors
	if err := c.ShouldBind(v); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}

		ts := c.Value("trans")
		trans := ts.(ut.Translator)
		for key, value := range errors.Translate(trans) {
			errs = append(errs, &ValidError{key, value})
		}
		return false, errs
	}
	return true, nil
}
