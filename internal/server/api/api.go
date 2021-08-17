package api

import (
	"fmt"
	"net/http"

	"goat/internal/model"
	"goat/internal/server/middleware"
	"goat/internal/svc"
	"goat/pkg/e"
	"goat/pkg/validatorer"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Base struct {
	Svc *svc.Service
	Log *zap.Logger
}

func New(store *model.Store, log *zap.Logger) Base {
	return Base{
		Svc: svc.New(store, log),
		Log: log.Named("api"),
	}
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OK 成功请求
func (b *Base) OK(c *gin.Context, arg ...interface{}) {
	resp := Response{Code: 0, Msg: "OK", Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	middleware.SetCode(c, resp.Code, resp.Msg)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

// Err 错误请求
func (b *Base) Err(c *gin.Context, httpCode int, err error, log bool, arg ...interface{}) {
	code, msg := e.DecodeErr(err)
	resp := Response{Code: code, Msg: msg, Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	middleware.SetCode(c, code, msg)
	c.AbortWithStatusJSON(httpCode, resp)
	if log {
		b.Log.Error(fmt.Sprintf("%+v", err))
	}
}

// HasErr 判断业务错误
func (b *Base) HasErr(c *gin.Context, err error) bool {
	if err != nil {
		b.Err(c, http.StatusInternalServerError, err, false)
		return true
	}
	return false
}

// HasErrL 判断业务错误，并打印错误
func (b *Base) HasErrL(c *gin.Context, err error) bool {
	if err != nil {
		b.Err(c, http.StatusInternalServerError, err, true)
		return true
	}
	return false
}

func (b *Base) ParseUri(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindUri(request); err != nil {
		return b.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseQuery(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindQuery(request); err != nil {
		return b.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseJSON(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		return b.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseForm(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindWith(request, binding.Form); err != nil {
		return b.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

// ValidatorData 参数校验
// hideDetails 可选择隐藏参数校验详细信息
func (b *Base) ValidatorData(c *gin.Context, err error, hideDetails bool) bool {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if eno, ok := err.(validator.ValidationErrors); ok {
			if !hideDetails {
				b.Err(c, http.StatusBadRequest, e.ErrValidation, false, validatorer.TranslateErrData(eno))
			} else {
				b.Err(c, http.StatusBadRequest, e.ErrValidation, false)
			}
			return false
		}
	}
	b.Err(c, http.StatusInternalServerError, e.ErrBind, false)
	return false
}
