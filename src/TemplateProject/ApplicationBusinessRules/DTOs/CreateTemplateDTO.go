package DTOs

type CreateTemplateDTO struct {
	Name string `json:"name" validate:"required"`
}
