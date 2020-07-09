package nullable

func PtrToStringOr(v *string, value string) string {
	if v == nil {
		return value
	}
	return *v
}

func PtrToString(v *string) string {
	return PtrToStringOr(v, "")
}

func StringToPtr(v string) *string {
	return &v
}
