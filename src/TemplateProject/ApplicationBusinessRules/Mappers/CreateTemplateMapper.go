package Mappers

import (
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
)

func CreateTemplateDtoToTemplateEntity(dto DTOs.CreateTemplateDTO) Entities.TemplateEntity {
	return Entities.TemplateEntity{
		Name: dto.Name,
	}
}
