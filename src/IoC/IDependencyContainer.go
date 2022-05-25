package IoC

import "TemplateSolution/src/TemplateProject"

type IDependencyContainer interface {
	NewTemplateContainer() TemplateProject.ITemplateContainer
}
