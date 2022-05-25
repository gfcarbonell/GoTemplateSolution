package GetTemplateByIdUseCase

import (
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/Ports/GetTemplateByIdPort"
)

type GetTemplateByIdInputPort struct {
	GetTemplateByIdPort.IGetTemplateByIdInputPort
	RequestData int64
	OutputPort  GetTemplateByIdPort.IGetTemplateByIdOutputPort
}

func NewGetTemplateByIdInputPort(requestData int64, outputPort GetTemplateByIdPort.IGetTemplateByIdOutputPort) GetTemplateByIdPort.IGetTemplateByIdInputPort {
	return &GetTemplateByIdInputPort{RequestData: requestData, OutputPort: outputPort}
}

func (iam *GetTemplateByIdInputPort) GetRequestData() int64 {
	return iam.RequestData
}

func (iam *GetTemplateByIdInputPort) GetOutputPort() GetTemplateByIdPort.IGetTemplateByIdOutputPort {
	return iam.OutputPort
}
