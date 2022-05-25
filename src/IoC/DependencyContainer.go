package IoC

import (
	"TemplateSolution/src/TemplateProject"
	"database/sql"
)

type DependencyContainer struct {
	IDependencyContainer
	Db *sql.DB
}

func NewDependencyContainer(db *sql.DB) IDependencyContainer {
	return &DependencyContainer{Db: db}
}

func (iam *DependencyContainer) NewTemplateContainer() TemplateProject.ITemplateContainer {
	return TemplateProject.NewTemplateContainer(iam.Db)
}
