package Presenters

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
)

type IGetTemplateByIdPresenter interface {
	Generic.IPresenter[DTOs.ReadTemplateDTO, DTOs.ReadTemplateDTO, error]
}
