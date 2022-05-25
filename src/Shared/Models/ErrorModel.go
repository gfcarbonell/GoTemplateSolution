package Models

type ErrorModel struct {
	error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorModel(code int, message string) *ErrorModel {
	return &ErrorModel{
		Code:    code,
		Message: message,
	}
}

func (e *ErrorModel) SetCode(code int) {
	e.Code = code
}

func (e *ErrorModel) GetCode() int {
	return e.Code
}

func (e *ErrorModel) SetMessage(message string) {
	e.Message = message
}

func (e *ErrorModel) GetMessage() string {
	return e.Message
}

func (e *ErrorModel) GetData() (int, string) {
	return e.Code, e.Message
}
