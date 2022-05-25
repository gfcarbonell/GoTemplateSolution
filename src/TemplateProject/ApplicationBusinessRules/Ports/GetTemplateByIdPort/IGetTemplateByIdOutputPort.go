package GetTemplateByIdPort

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
)

type IGetTemplateByIdOutputPort interface {
	Generic.IOutputPort[DTOs.ReadTemplateDTO, error]
}
