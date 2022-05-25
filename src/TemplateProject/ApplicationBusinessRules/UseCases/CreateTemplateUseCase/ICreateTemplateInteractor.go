package CreateTemplateUseCase

type ICreateTemplateInteractor interface {
	Handle(response *CreateTemplateInputPort)
}
