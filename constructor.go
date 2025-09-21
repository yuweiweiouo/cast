package cast

type Constructor[T any, R any] struct {
	caster Caster[T, R]
}

func (c Constructor[T, R]) FromValue(t T) *scalarCastDelegator[T, R] {
	return &scalarCastDelegator[T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carrier:       Carry(&t),
	}
}

func (c Constructor[T, R]) FromPointer(pt *T) *scalarCastDelegator[T, R] {
	return &scalarCastDelegator[T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carrier:       Carry(pt),
	}
}

func (c Constructor[T, R]) FromValues(ts []T) *sliceCastDelegator[T, R] {
	d := &sliceCastDelegator[T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carriers:      make([]Carrier[T], len(ts)),
	}
	for i, t := range ts {
		d.carriers[i] = Carry(&t)
	}
	return d
}

func (c Constructor[T, R]) FromPointers(pts []*T) *sliceCastDelegator[T, R] {
	d := &sliceCastDelegator[T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carriers:      make([]Carrier[T], len(pts)),
	}
	for i, ptr := range pts {
		d.carriers[i] = Carry(ptr)
	}
	return d
}

func (c Constructor[T, R]) FromStringMap(tm map[string]T) *mapCastDelegator[string, T, R] {
	d := &mapCastDelegator[string, T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carrierMap:    make(map[string]Carrier[T]),
	}
	for k, t := range tm {
		d.carrierMap[k] = Carry(&t)
	}
	return d
}

func (c Constructor[T, R]) FromIntMap(tm map[int]T) *mapCastDelegator[int, T, R] {
	d := &mapCastDelegator[int, T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carrierMap:    make(map[int]Carrier[T]),
	}
	for k, t := range tm {
		d.carrierMap[k] = Carry(&t)
	}
	return d
}

func (c Constructor[T, R]) FromInt64Map(tm map[int64]T) *mapCastDelegator[int64, T, R] {
	d := &mapCastDelegator[int64, T, R]{
		castDelegator: castDelegator[T, R]{caster: c.caster},
		carrierMap:    make(map[int64]Carrier[T]),
	}
	for k, t := range tm {
		d.carrierMap[k] = Carry(&t)
	}
	return d
}

// 偷懶用
func (c Constructor[T, R]) WithValue(t T) R {
	return c.FromValue(t).ToValue()
}

func (c Constructor[T, R]) WithPointer(pt *T) *R {
	return c.FromPointer(pt).ToPointer()
}

func (c Constructor[T, R]) WithValues(ts []T) []R {
	return c.FromValues(ts).ToValues()
}

func (c Constructor[T, R]) WithPointers(pts []*T) []*R {
	return c.FromPointers(pts).ToPointers()
}
