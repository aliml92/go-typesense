package typesense

func Bool(v bool) *bool {
	return &v
}

func String(v string) *string {
	return &v
}

func Int64(v int) *int64 {
	n := int64(v)
	return &n
}

func Int(v int) *int {
	return &v
}
