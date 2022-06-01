package TemplateProject

import "TemplateSolution/src/TemplateProject/InterfaceAdapters/Controllers"

type ITemplateContainer interface {
	NewTemplateController() Controllers.TemplateController
}
