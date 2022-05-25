package CreateTemplateUseCase

import (
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/Ports/CreateTemplatePort"
)

type CreateTemplateInputPort struct {
	CreateTemplatePort.ICreateTemplateInputPort
	RequestData DTOs.CreateTemplateDTO
	OutputPort  CreateTemplatePort.ICreateTemplateOutputPort
}

func NewCreateTemplateInputPort(requestData DTOs.CreateTemplateDTO, outputPort CreateTemplatePort.ICreateTemplateOutputPort) CreateTemplatePort.ICreateTemplateInputPort {
	return &CreateTemplateInputPort{RequestData: requestData, OutputPort: outputPort}
}

func (iam *CreateTemplateInputPort) GetRequestData() DTOs.CreateTemplateDTO {
	return iam.RequestData
}

func (iam *CreateTemplateInputPort) GetOutputPort() CreateTemplatePort.ICreateTemplateOutputPort {
	return iam.OutputPort
}
