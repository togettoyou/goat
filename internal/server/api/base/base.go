package base

import (
	"fmt"
	"net/http"

	"goat/pkg/e"
	"goat/pkg/validatorer"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Api struct {
	Log *zap.Logger
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OK 成功请求
func (api *Api) OK(c *gin.Context, arg ...interface{}) {
	resp := response{Code: 0, Msg: "OK", Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

// Err 错误请求
func (api *Api) Err(c *gin.Context, err error, log bool, arg ...interface{}) {
	code, msg := e.DecodeErr(err)
	resp := response{Code: code, Msg: msg, Data: map[string]interface{}{}}
	if len(arg) > 0 {
		resp.Data = arg[0]
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, resp)
	if log {
		api.Log.Error(fmt.Sprintf("%+v", err))
	}
}

// HasErr 判断业务错误
func (api *Api) HasErr(c *gin.Context, err error) bool {
	if err != nil {
		api.Err(c, err, false)
		return true
	}
	return false
}

// HasErrL 判断业务错误，并打印错误
func (api *Api) HasErrL(c *gin.Context, err error) bool {
	if err != nil {
		api.Err(c, err, true)
		return true
	}
	return false
}

func (api *Api) ParseUri(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindUri(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], false)
	}
	return true
}

func (api *Api) ParseQuery(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindQuery(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], false)
	}
	return true
}

func (api *Api) ParseJSON(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], false)
	}
	return true
}

func (api *Api) ParseForm(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindWith(request, binding.Form); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], false)
	}
	return true
}

func (api *Api) ParseUriL(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindUri(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], true)
	}
	return true
}

func (api *Api) ParseQueryL(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindQuery(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], true)
	}
	return true
}

func (api *Api) ParseJSONL(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], true)
	}
	return true
}

func (api *Api) ParseFormL(c *gin.Context, request interface{}, hideDetails ...bool) bool {
	if err := c.ShouldBindWith(request, binding.Form); err != nil {
		return api.ValidatorData(c, err, len(hideDetails) > 0 && hideDetails[0], true)
	}
	return true
}

// ValidatorData 参数校验
// hideDetails 可选择隐藏参数校验详细信息
func (api *Api) ValidatorData(c *gin.Context, err error, hideDetails bool, log bool) bool {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if eno, ok := err.(validator.ValidationErrors); ok {
			if !hideDetails {
				api.Err(c, e.ErrValidation, log, validatorer.TranslateErrData(eno))
			} else {
				api.Err(c, e.ErrValidation, log)
			}
			return false
		}
	}
	api.Err(c, e.ErrBind, log)
	return false
}
