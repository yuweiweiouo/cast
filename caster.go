package cast

type Caster[T any, R any] interface {
	Cast(ca Carrier[T]) (cb Carrier[R])
}

type CastFunc[T any, R any] func(ca Carrier[T]) (cb Carrier[R])

func (cf CastFunc[T, R]) Cast(ca Carrier[T]) (cb Carrier[R]) {
	return cf(ca)
}
