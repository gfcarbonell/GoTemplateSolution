package CreateTemplateUseCase

import (
	"TemplateSolution/src/Shared/Constants"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/Mappers"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories"
	"gopkg.in/go-playground/validator.v9"
)

type CreateTemplateInteractor struct {
	ICreateTemplateInteractor
	CreateTemplateRepository Repositories.ICreateTemplateRepository
}

func NewCreateTemplateInteractor(createTemplateRepository Repositories.ICreateTemplateRepository) ICreateTemplateInteractor {
	return &CreateTemplateInteractor{CreateTemplateRepository: createTemplateRepository}
}

func (iam *CreateTemplateInteractor) Handle(request *CreateTemplateInputPort) {
	valid := validator.New()
	err := valid.Struct(request.GetRequestData())

	if err != nil {
		//request.GetOutputPort().Handle(-1, err.(validator.ValidationErrors))
		request.GetOutputPort().Handle(-1, Constants.ErrValidation)
		return
	}

	var entity = Mappers.CreateTemplateDtoToTemplateEntity(request.GetRequestData())
	id, err := iam.CreateTemplateRepository.Handle(entity)

	if err != nil {
		request.GetOutputPort().Handle(-1, err)
		return
	}

	request.GetOutputPort().Handle(*id, nil)
}
