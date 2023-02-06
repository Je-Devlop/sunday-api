package router

import (
	"Je-Devlop/sunday-api/sunday"

	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *MyContext) JSON(status int, v interface{}) {
	c.Context.JSON(status, v)
}

func NewGinHandler(handler func(sunday.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()

	return &MyRouter{r}
}

func (group *MyRouter) POST(relativePath string, handler func(sunday.Context)) {
	group.Engine.POST(relativePath, NewGinHandler(handler))
}
