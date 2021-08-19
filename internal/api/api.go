package api

import (
	"fmt"
	"net/http"

	"goat-layout/internal/model"
	"goat-layout/internal/server/middleware"
	"goat-layout/internal/svc"
	"goat-layout/pkg/e"
	logpkg "goat-layout/pkg/log"
	"goat-layout/pkg/validatorer"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Base 每个api对象都需要内嵌该结构体
type Base struct {
	log   *zap.Logger
	store *model.Store
	c     *gin.Context

	Log *zap.Logger
}

func New(store *model.Store, log *zap.Logger) Base {
	if log == nil {
		log = logpkg.New("").L()
	}
	return Base{
		log:   log,
		store: store,
		Log:   log.Named("api"),
	}
}

// MakeContext 设置http上下文
func (b *Base) MakeContext(c *gin.Context) *Base {
	b.c = c
	return b
}

// MakeService 设置业务对象
func (b *Base) MakeService(svc *svc.Service) *Base {
	if svc == nil {
		return b
	}
	svc.New(b.store, b.log)
	return b
}

func (b *Base) Named(name string) *Base {
	if b.log == nil {
		b.log = logpkg.New("").L()
	}
	b.log = b.log.Named(name)
	b.Log = b.log.Named("api")
	return b
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OK 成功请求
func (b *Base) OK(arg ...interface{}) {
	resp := Response{Code: 0, Msg: "OK", Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	middleware.SetCode(b.c, resp.Code, resp.Msg)
	b.c.AbortWithStatusJSON(http.StatusOK, resp)
}

// Err 错误请求
func (b *Base) Err(httpCode int, err error, log bool, arg ...interface{}) {
	code, msg := e.DecodeErr(err)
	resp := Response{Code: code, Msg: msg, Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	middleware.SetCode(b.c, code, msg)
	b.c.AbortWithStatusJSON(httpCode, resp)
	if log {
		b.log.Error(fmt.Sprintf("%+v", err))
	}
}

// HasErr 判断业务错误
func (b *Base) HasErr(err error) bool {
	if err != nil {
		b.Err(http.StatusInternalServerError, err, false)
		return true
	}
	return false
}

// HasErrL 判断业务错误，并打印错误
func (b *Base) HasErrL(err error) bool {
	if err != nil {
		b.Err(http.StatusInternalServerError, err, true)
		return true
	}
	return false
}

func (b *Base) ParseUri(request interface{}, hideDetails ...bool) bool {
	if err := b.c.ShouldBindUri(request); err != nil {
		return b.ValidatorData(err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseQuery(request interface{}, hideDetails ...bool) bool {
	if err := b.c.ShouldBindQuery(request); err != nil {
		return b.ValidatorData(err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseJSON(request interface{}, hideDetails ...bool) bool {
	if err := b.c.ShouldBindJSON(request); err != nil {
		return b.ValidatorData(err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

func (b *Base) ParseForm(request interface{}, hideDetails ...bool) bool {
	if err := b.c.ShouldBindWith(request, binding.Form); err != nil {
		return b.ValidatorData(err, len(hideDetails) > 0 && hideDetails[0])
	}
	return true
}

// ValidatorData 参数校验
// hideDetails 可选择隐藏参数校验详细信息
func (b *Base) ValidatorData(err error, hideDetails bool) bool {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if eno, ok := err.(validator.ValidationErrors); ok {
			if !hideDetails {
				b.Err(http.StatusBadRequest, e.ErrValidation, false, validatorer.TranslateErrData(eno))
			} else {
				b.Err(http.StatusBadRequest, e.ErrValidation, false)
			}
			return false
		}
	}
	b.Err(http.StatusInternalServerError, e.ErrBind, false)
	return false
}
