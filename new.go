package generic

func New[Struct, T any](iteration func() T) (out Struct) {
	Modify(&out, func(v *T) {
		*v = iteration()
	})
	return
}

func NewArray[T any](count uint, iteration func() T) []T {
	out := make([]T, count)
	for i := range out {
		out[i] = iteration()
	}
	return out
}
