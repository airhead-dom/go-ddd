package usecase

type UseCase[T any] interface {
	Do(param T) (any, error)
}
