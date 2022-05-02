package vector

// Iterator allows iteration over values of a list (only once)
type Iterator func() any

func (f Iterator) ToSlice() []any {
	s := make([]any, 0)
	for {
		v := f()
		if v == nil {
			break
		}
		s = append(s, v)
	}
	return s
}
