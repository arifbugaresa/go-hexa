package health

import (
	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Health(context echo.Context) (err error) {
	return context.NoContent(200)
}
