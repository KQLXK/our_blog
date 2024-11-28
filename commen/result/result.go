package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var (
	RegisterErrStatus          = newstatus(http.StatusBadRequest, 40001, "注册发生错误")
	UsernameExistErrStatus     = newstatus(http.StatusBadRequest, 40002, "用户名已存在")
	EmailExistErrStatus        = newstatus(http.StatusBadRequest, 40003, "邮箱已被注册")
	ParseTokenErrStatus        = newstatus(http.StatusBadRequest, 40004, "token格式有误，请检查你的操作")
	GetReqErrStatus            = newstatus(http.StatusBadRequest, 40005, "无法获取请求参数")
	UsernameNotExsitsErrStatus = newstatus(http.StatusBadRequest, 40006, "用户名不存在")
	PasswordWrongErr           = newstatus(http.StatusBadRequest, 40007, "密码错误")

	//401未被授权的
	GetTokenErrStatus = newstatus(http.StatusUnauthorized, 40101, "请先进行登录")
	TokenExiredStatus = newstatus(http.StatusUnauthorized, 40102, "token过期，请重新登录")

	//500服务器内部错误
	ServerErrStatus = newstatus(http.StatusInternalServerError, 50000, "服务器内部错误")
)

type Status struct {
	HTTPcode   int
	StatusCode int
	Message    string
}

func (s *Status) httpcode() int {
	return s.HTTPcode
}

func (s *Status) statuscode() int {
	return s.StatusCode
}

func (s *Status) message() string {
	return s.Message
}

func newstatus(httpcode int, statuscode int, message string) Status {
	return Status{
		HTTPcode:   httpcode,
		StatusCode: httpcode,
		Message:    message,
	}
}

type R map[string]interface{}

func (r R) ToMap(s interface{}) {
	// 使用反射获取指向的值
	val := reflect.ValueOf(s)

	// 如果是指针，获取指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 迭代结构体字段
	typeOfS := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typeOfS.Field(i)
		value := val.Field(i)

		// 将字段的名称和值添加到 r 中
		r[field.Name] = value.Interface()
	}
}

func Sucess(c *gin.Context, r R) {
	h := gin.H{
		"status":  0,
		"message": "sucess",
	}
	h["data"] = r
	c.JSON(http.StatusOK, h)
}

func Error(c *gin.Context, s Status) {
	c.JSON(s.httpcode(), gin.H{
		"status":  s.StatusCode,
		"message": s.Message,
	})
}
