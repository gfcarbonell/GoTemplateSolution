package Contexts

type ControllerContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}
