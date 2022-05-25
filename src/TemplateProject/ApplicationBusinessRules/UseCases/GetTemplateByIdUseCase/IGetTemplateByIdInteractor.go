package GetTemplateByIdUseCase

type IGetTemplateByIdInteractor interface {
	Handle(response *GetTemplateByIdInputPort)
}
