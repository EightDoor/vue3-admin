package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strings"
	"zhoukai/configure"
	"zhoukai/model"
	"zhoukai/utils/db"
)

type Paging struct {
	Page int
	Size int
}
// 定义一个全局翻译器T
var trans ut.Translator
// 公共请求返回
func R(database interface{},result *gorm.DB, c *gin.Context) {
	err := result.Error
	// 总数
	//dates, _ := result.Rows()
	pageSize, page := UtilsDB.PageIngService(c)
	responseData := model.ResponseData{
		Page:     page,
		PageSize: pageSize,
		List:     database,
	}
	response := model.Response{
		Code:    configure.RequestOtherError,
		Message: err,
		Data:    responseData,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,response)
	}else {
		response.Code = configure.RequestSuccess
		response.Message = "success"
		c.JSON(http.StatusOK, response)
	}
}
// 参数校验返回
func Verify(data interface{}, c *gin.Context) error {
	//validate := validator.New()
	//// 自定义校验返回的参数格式
	//InitTrans("zh", validate)
	//err := validate.Struct(data)
	err := c.ShouldBind(data)
	if err != nil {
		//var errValue interface{}
		//// 获取validator.ValidationErrors类型的errors
		//errs, ok := err.(validator.ValidationErrors)
		//
		//if !ok {
		//	errValue = err.Error()
		//	// 非validator.ValidationErrors类型错误直接返回
		//}else {
		//	errValue = removeTopStruct(errs.Translate(trans))
		//}
		response := model.Verification{
			Message: err,
			Code: configure.RequestKeyNotFound,
		}
		c.JSON(http.StatusOK, response)
		return err
	}else {
		return nil
	}
}

func InitTrans(locale string, validateS *validator.Validate)(err error)  {
	// 注册一个获取json tag的自定义方法
	validateS.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"),",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// 中文翻译器
	zhT := zh.New()
	enT := en.New()
	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)
	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	trans, ok := uni.GetTranslator(locale)
	if ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(validateS, trans)
		break
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(validateS, trans)
	}
	if err != nil {
		return err
	}
	return
}

// removeTopStruct 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}