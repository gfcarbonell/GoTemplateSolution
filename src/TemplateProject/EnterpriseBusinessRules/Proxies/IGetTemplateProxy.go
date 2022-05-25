package Proxies

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Models"
)

type IGetTemplateByIdProxy interface {
	Generic.IService[int64, *Models.TemplateModel, error]
}
