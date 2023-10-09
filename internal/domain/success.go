package domain

type SuccessResult[T any] struct {
	Data T `json:"data"`
}
