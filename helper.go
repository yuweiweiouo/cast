package cast

func empty[T any]() T {
	var zero T
	return zero
}

func isNil[T any](pt *T) bool {
	return pt == nil || pt == (*T)(nil)
}
