package valid

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)


func InitTrans() (trans ut.Translator) {
	v := binding.Validator.Engine().(*validator.Validate)
	zhT := zh.New()
	uni := ut.New(zhT)
	trans, _ = uni.GetTranslator("zh")
	zh2.RegisterDefaultTranslations(v, trans)
	return
}

