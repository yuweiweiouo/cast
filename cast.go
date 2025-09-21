package cast

func Use[T any, R any](castFunc func(ca Carrier[T]) (cb Carrier[R])) Constructor[T, R] {
	return UseCaster(CastFunc[T, R](castFunc))
}

func UseCaster[T any, R any](caster Caster[T, R]) Constructor[T, R] {
	return Constructor[T, R]{
		caster: caster,
	}
}
