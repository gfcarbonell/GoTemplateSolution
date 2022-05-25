package Controllers

import (
	"github.com/labstack/echo/v4"
)

type ITemplateController interface {
	CreateTemplate(ctx echo.Context) error
	GetTemplateByIdController(ctx echo.Context) error
}
