package usecase

type UseCases interface {
	Calculate(data string) (int, error)
}
