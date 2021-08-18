//go:build docs
// +build docs

package router

import (
	_ "goat-layout/docs"

	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swag = swagger.WrapHandler(swaggerFiles.Handler)
}
