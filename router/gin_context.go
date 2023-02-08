package router

import (
	"Je-Devlop/sunday-api/sunday"

	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{Context: c}
}

func NewGinHandler(handler func(sunday.FrameworkContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewGinContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()

	return &MyRouter{r}
}

func (group *MyRouter) POST(relativePath string, handler func(sunday.FrameworkContext)) {
	group.Engine.POST(relativePath, NewGinHandler(handler))
}

func (group *MyRouter) GET(relativePath string, handler func(sunday.FrameworkContext)) {
	group.Engine.GET(relativePath, NewGinHandler(handler))
}

// func healthz(e FrameworkRouter) {
// 	e.GET("/healthz", func(c NewGinContext) {
// 		c.Status(http.StatusOK)
// 	})
// }
