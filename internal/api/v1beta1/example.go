package v1beta1

import (
	"errors"

	"goat-layout/internal/api"
	"goat-layout/pkg/e"

	"github.com/gin-gonic/gin"
)

type Example struct {
	api.Base
}

// Get
// @Tags example
// @Summary Get请求
// @Produce json
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example [get]
func (g Example) Get(c *gin.Context) {
	g.Named("test").MakeContext(c)
	g.Log.Debug("Get请求")
	g.Log.Warn("Get请求")
	g.Log.Info("Get请求")
	g.Log.Error("Get请求")
	g.OK("打印日志")
}

// Err
// @Tags example
// @Summary Err请求
// @Produce json
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example/err [get]
func (g Example) Err(c *gin.Context) {
	g.MakeContext(c)

	// 这里模拟业务层处理业务后返回了一个err，注：err应该在svc层包装
	// err := e.ErrNotLogin // 返回的错误没有包装 ❌
	// err := e.Wrap(e.ErrNotLogin) // 返回的错误使用Wrap包装 ✔
	err := e.New(e.ErrNotLogin, errors.New("第三方错误或系统错误")) // 返回的错误是使用New创建的 ✔

	// 判断错误，g.HasErr不打印错误消息，g.HasErrL会打印错误消息
	if g.HasErrL(err) {
		return
	}
	g.OK(c)
}

type UriArgs struct {
	ID uint `json:"id" uri:"id" binding:"required,min=10"`
}

// Uri
// @Tags example
// @Summary uri参数请求
// @Description 路径参数，匹配 /uri/{id}
// @Produce json
// @Param id path int false "id值"
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example/uri/{id} [get]
func (g Example) Uri(c *gin.Context) {
	g.MakeContext(c)
	//id := c.Param("id")
	var args UriArgs
	if !g.ParseUri(&args) {
		return
	}
	g.OK(args)
}

type QueryArgs struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

// Query
// @Tags example
// @Summary query参数查询
// @Description 查询参数，匹配 query?id=xxx
// @Produce json
// @Param email query string true "邮箱"
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example/query [get]
func (g Example) Query(c *gin.Context) {
	g.MakeContext(c)
	//email := c.Query("email")
	var args QueryArgs
	if !g.ParseQuery(&args) {
		return
	}
	g.OK(args)
}

type FormArgs struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

// FormData
// @Tags example
// @Summary form表单请求
// @Description 处理application/x-www-form-urlencoded类型的POST请求
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "邮箱"
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example/form [post]
func (g Example) FormData(c *gin.Context) {
	g.MakeContext(c)
	//email := c.PostForm("email")
	var args FormArgs
	if !g.ParseForm(&args) {
		return
	}
	g.OK(args)
}

type JSONBody struct {
	Email    string `json:"email" binding:"required,email" example:"admin@qq.com"`
	Username string `json:"username" binding:"required,checkUsername" example:"admin"`
}

// JSON
// @Tags example
// @Summary JSON参数请求
// @Description 邮箱、用户名校验
// @Produce  json
// @Param data body JSONBody true "测试请求json参数"
// @Success 200 {object} api.Response
// @Router /api/v1beta1/example/json [post]
func (g Example) JSON(c *gin.Context) {
	var body JSONBody
	if !g.MakeContext(c).ParseJSON(&body) {
		return
	}
	g.OK(body)
}
