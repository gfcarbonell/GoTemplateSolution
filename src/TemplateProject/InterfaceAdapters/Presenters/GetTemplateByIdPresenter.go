package Presenters

import "TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"

type GetTemplateByIdPresenter struct {
	IGetTemplateByIdPresenter
	content DTOs.ReadTemplateDTO
	err     error
}

func NewGetTemplateByIdPresenter() IGetTemplateByIdPresenter {
	return &GetTemplateByIdPresenter{}
}

func (iam *GetTemplateByIdPresenter) Handle(response DTOs.ReadTemplateDTO, err error) {
	iam.content = response
	iam.err = err
}

func (iam *GetTemplateByIdPresenter) GetContent() DTOs.ReadTemplateDTO {
	return iam.content
}

func (iam *GetTemplateByIdPresenter) GetError() error {
	return iam.err
}
