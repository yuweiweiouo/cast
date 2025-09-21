package cast

func Carry[T any](t *T) Carrier[T] {
	return Carrier[T]{item: t}
}

func CarryValue[T any](t T) Carrier[T] {
	return Carrier[T]{item: &t}
}

// Carrier 存放 T型別 的 pointer
type Carrier[T any] struct {
	item *T
}

func (c Carrier[T]) IsCarrying() bool {
	return !isNil(c.item)
}

func (c Carrier[T]) Get() *T {
	return c.item
}

func (c Carrier[T]) GetVal() T {
	if c.IsCarrying() {
		return *c.item
	}
	return empty[T]()
}
