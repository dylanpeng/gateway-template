package router

import (
	"github.com/dylanpeng/gateway-template/common"
	ctrl "github.com/dylanpeng/gateway-template/common/control"
	"github.com/dylanpeng/gateway-template/common/middleware"
	"github.com/gin-gonic/gin"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	return common.GetTraceId(ctx)
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)
	app.Use(middleware.Trace)
}
