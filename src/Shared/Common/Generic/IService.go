package Generic

type IService[TInputModel any, TOutputModel any, TError any] interface {
	Handle(input TInputModel) (TOutputModel, TError)
}
