package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
// 公共请求返回(普通返回)
func R(database interface{},result *gorm.DB, c *gin.Context) {
	err := result.Error
	rowsAffected := result.RowsAffected
	pageSize, page := UtilsDB.PageIngService(c)
	total := UtilsDB.GetTotal(c)
	// 反射获取类型
	typeOfA := reflect.TypeOf(database)
	var responseData interface{}
	responseData = model.ResponseDataSing{
		List: database,
	}
	 typeOfB := typeOfA.Kind() == reflect.Slice
	if typeOfB {
		responseData = model.ResponseData{
			List: database,
			PageSize: pageSize,
			Page: page,
			Total: total,
		}
	}

	response := model.Response{
		Code:    configure.RequestOtherError,
		Message: err,
		Data:    responseData,
	}
	if err != nil {
		c.JSON(http.StatusOK,response)
	}else if rowsAffected == 0 && !typeOfB {
		response.Message = "操作影响行数为 -> 0"
		c.JSON(http.StatusOK,response)
	} else {
		response.Code = configure.RequestSuccess
		response.Message = "success"
		c.JSON(http.StatusOK, response)
	}
}
// 参数校验返回
func Verify(data interface{}, c *gin.Context) error {
	InitTrans("zh")
	err := c.ShouldBind(data)
	if err != nil {
		var msg interface{}
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			msg = err.Error()
		}else {
			msg = removeTopStruct(errs.Translate(trans))
		}
		response := model.Verification{
			Message: msg,
			Code: configure.RequestKeyNotFound,
		}
		c.JSON(http.StatusOK, response)
		return err
	}else {
		return nil
	}
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {

	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"),",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
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