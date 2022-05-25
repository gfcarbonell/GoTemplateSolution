package Bases

import "github.com/labstack/echo/v4"

type IBaseController interface {
	Handle(ctx echo.Context) error
}
