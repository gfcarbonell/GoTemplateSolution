package Repositories

import (
	"TemplateSolution/src/Shared/Common/Generic"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
)

//go:generate mockery --name=ICreateTemplateRepository --filename=CreateTemplateRepositoryMock.go --outpkg=mockRepositories
type ICreateTemplateRepository interface {
	Generic.IRepository[Entities.TemplateEntity, *int64, error]
}
