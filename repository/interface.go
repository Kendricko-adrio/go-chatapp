package repository

type IRepository[T any, V any] interface {
	FindAll() []T
	FindById(V) T
}
