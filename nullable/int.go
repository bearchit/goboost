package nullable

func PtrToIntOr(v *int, value int) int {
	if v == nil {
		return value
	}
	return *v
}

func PtrToInt(v *int) int {
	return PtrToIntOr(v, 0)
}

func IntToPtr(v int) *int {
	return &v
}

func DefaultIfNilInt(v *int, dv int) int {
	if v != nil && *v > 0 {
		return *v
	}
	return dv
}
