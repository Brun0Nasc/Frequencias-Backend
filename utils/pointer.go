package utils

func GetInt64Pointer(value int) *int64 {
	v := int64(value)
	return &v
}
