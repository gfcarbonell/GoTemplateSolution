package Constants

import (
	"TemplateSolution/src/Shared/Models"
)

var (
	ErrAuth              = &Models.ErrorModel{Code: 0, Message: "invalid token"}
	ErrGenerateInvalidId = &Models.ErrorModel{Code: 1, Message: "Generate invalid Id"}
	ErrNotFound          = &Models.ErrorModel{Code: 2, Message: "not found"}
	ErrDuplicate         = &Models.ErrorModel{Code: 3, Message: "duplicate"}
	ErrInternalServer    = &Models.ErrorModel{Code: 4, Message: "Internal Server"}
	ErrBadRequest        = &Models.ErrorModel{Code: 4, Message: "Bad Request"}
	ErrValidation        = &Models.ErrorModel{Code: 4, Message: "Validation Fields Error"}
)
