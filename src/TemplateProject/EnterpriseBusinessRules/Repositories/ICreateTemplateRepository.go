package Repositories

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
)

type ICreateTemplateRepository interface {
	Generic.IRepository[*Entities.TemplateEntity, *int64, error]
}
