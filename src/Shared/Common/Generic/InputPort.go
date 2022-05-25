package Generic

type IInputPort[InteractorRequestType any, InteractorResponseType any] interface {
	GetRequestData() InteractorRequestType
	GetOutputPort() InteractorResponseType
}
