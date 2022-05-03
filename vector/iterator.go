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

func (f Iterator) ToInvertedSlice() []any {
	s := f.ToSlice()
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}
