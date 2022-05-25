package Controllers

type TemplateController struct {
	CreateTemplateController interface {
		ICreateTemplateController
	}
	GetTemplateByIdController interface {
		IGetTemplateByIdController
	}
}
