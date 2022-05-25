package Generic

type IOutputPort[InteractorResponseType any, TError any] interface {
	Handle(response InteractorResponseType, err TError)
	GetError() TError
}
