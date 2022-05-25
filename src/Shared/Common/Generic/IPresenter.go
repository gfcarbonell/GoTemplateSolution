package Generic

type IPresenter[ResponseType any, FormatType any, TError any] interface {
	IOutputPort[ResponseType, TError]
	GetContent() FormatType
}
