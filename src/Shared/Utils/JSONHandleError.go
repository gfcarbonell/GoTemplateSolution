package Utils

import (
	"TemplateSolution/src/Shared/Contexts"
	"TemplateSolution/src/Shared/Models"
	"errors"
	"net/http"
)

func JSONHandleError(statusCode int, err error, ctx Contexts.ControllerContext) error {
	var errorCustom Models.ErrorModel
	if errors.As(err, &errorCustom) {
		return ctx.JSON(statusCode, errorCustom)
	} else {
		errorCustom.SetCode(-1)
		errorCustom.SetMessage("Oop! Error internal server!")

		return ctx.JSON(http.StatusInternalServerError, errorCustom)
	}
}
