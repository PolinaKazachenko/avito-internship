package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterHandlers(g *gin.RouterGroup)
}

func RegisterHandlers(router *gin.Engine, handlers ...Handler) {
	rg := router.Group("/")
	for _, handler := range handlers {
		handler.RegisterHandlers(rg)
	}
}
