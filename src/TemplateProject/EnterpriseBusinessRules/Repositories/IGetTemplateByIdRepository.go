package Repositories

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
)

type IGetTemplateByIdRepository interface {
	Generic.IRepository[int64, *Entities.TemplateEntity, error]
}
