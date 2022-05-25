package CreateTemplatePort

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
)

type ICreateTemplateInputPort interface {
	Generic.IInputPort[DTOs.CreateTemplateDTO, ICreateTemplateOutputPort]
}
