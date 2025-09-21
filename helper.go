package cast

func handle[T any, R any](t Carrier[T], caster Caster[T, R]) Carrier[R] {
	if t.IsCarrying() {
		return caster.Cast(t)
	}
	return Carry[R](nil)
}

func empty[T any]() T {
	var zero T
	return zero
}

func isNil[T any](pt *T) bool {
	return pt == nil || pt == (*T)(nil)
}
