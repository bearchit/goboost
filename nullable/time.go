package nullable

import "time"

func PtrToTimeOr(v *time.Time, value time.Time) time.Time {
	if v == nil {
		return value
	}
	return *v
}

func PtrToTime(v *time.Time) time.Time {
	return PtrToTimeOr(v, time.Time{})
}

func TimeToPtr(v time.Time) *time.Time {
	return &v
}
