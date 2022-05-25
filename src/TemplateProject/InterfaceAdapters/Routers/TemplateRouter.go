package Routers

import (
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Controllers"
	"github.com/labstack/echo/v4"
)

func NewTemplateRouter(e *echo.Group, c Controllers.TemplateController) *echo.Group {
	e.POST("/", func(context echo.Context) error {
		return c.CreateTemplateController.Handle(context)
	})

	e.GET("/:id", func(context echo.Context) error {
		return c.GetTemplateByIdController.Handle(context)
	})

	return e
}
