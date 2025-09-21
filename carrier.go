package cast

func Carry[T any](t *T) Carrier[T] {
	return Carrier[T]{object: t}
}

func CarryValue[T any](t T) Carrier[T] {
	return Carrier[T]{object: &t}
}

// Carrier 存放 T型別 的 pointer
type Carrier[T any] struct {
	object *T
}

func (c Carrier[T]) IsCarrying() bool {
	return !isNil(c.object)
}

func (c Carrier[T]) Get() *T {
	return c.object
}

func (c Carrier[T]) GetVal() T {
	if c.IsCarrying() {
		return *c.object
	}
	return empty[T]()
}
