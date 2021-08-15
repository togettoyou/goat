//+build docs

package router

import (
	_ "goat/cmd/server/docs"

	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swag = swagger.WrapHandler(swaggerFiles.Handler)
}
