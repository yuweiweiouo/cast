package cast

type castDelegator[T any, R any] struct {
	caster Caster[T, R]
}

func (d *castDelegator[T, R]) handle(t Carrier[T]) Carrier[R] {
	if t.IsCarrying() {
		return d.caster.Cast(t)
	}
	return Carry[R](nil)
}

type scalarCastDelegator[T any, R any] struct {
	castDelegator[T, R]
	carrier Carrier[T]
}

func (d *scalarCastDelegator[T, R]) ToValue() R {
	return d.handle(d.carrier).GetVal()
}

func (d *scalarCastDelegator[T, R]) ToPointer() *R {
	return d.handle(d.carrier).Get()
}

type sliceCastDelegator[T any, R any] struct {
	castDelegator[T, R]
	carriers []Carrier[T]
}

func (d *sliceCastDelegator[T, R]) ToValues() []R {
	rs := make([]R, len(d.carriers))
	for i, carrier := range d.carriers {
		rs[i] = d.handle(carrier).GetVal()
	}
	return rs
}

func (d *sliceCastDelegator[T, R]) ToPointers() []*R {
	rs := make([]*R, len(d.carriers))
	for i, carrier := range d.carriers {
		rs[i] = d.handle(carrier).Get()
	}
	return rs
}
