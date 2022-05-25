package Generic

type IRepository[TInputEntity any, TOutputEntity any, TError any] interface {
	Handle(entity TInputEntity) (TOutputEntity, TError)
}
