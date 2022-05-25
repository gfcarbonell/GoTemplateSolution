package GetTemplateByIdUseCase

import (
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Proxies"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories"
)

type GetTemplateByIdInteractor struct {
	IGetTemplateByIdInteractor
	GetTemplateByIdProxy      Proxies.IGetTemplateByIdProxy
	GetTemplateByIdRepository Repositories.IGetTemplateByIdRepository
}

func NewGetTemplateByIdInteractor(getTemplateByIdProxy Proxies.IGetTemplateByIdProxy, getTemplateByIdRepository Repositories.IGetTemplateByIdRepository) IGetTemplateByIdInteractor {
	return &GetTemplateByIdInteractor{
		GetTemplateByIdProxy:      getTemplateByIdProxy,
		GetTemplateByIdRepository: getTemplateByIdRepository,
	}
}

func (iam *GetTemplateByIdInteractor) Handle(request *GetTemplateByIdInputPort) {
	var dto DTOs.ReadTemplateDTO

	entity, err := iam.GetTemplateByIdRepository.Handle(request.GetRequestData())

	if err != nil {
		request.GetOutputPort().Handle(dto, err)
		return
	}

	model, _ := iam.GetTemplateByIdProxy.Handle(request.GetRequestData())

	if err != nil {
		request.GetOutputPort().Handle(dto, err)
		return
	}

	dto.Id = entity.Id
	dto.Name = entity.Name
	dto.Description = model.Description

	request.GetOutputPort().Handle(dto, nil)
}
