package xis

// InStr 检查string值在一个string slice中是否存在
func InStr(s string, ss []string) bool {
	for k := range ss {
		if s == ss[k] {
			return true
		}
	}
	return false
}

// InInt 检查int值在一个int slice中是否存在
func InInt(i int, is []int) bool {
	for k := range is {
		if i == is[k] {
			return true
		}
	}
	return false
}

// InInt64 检查int64值在一个int64 slice中是否存在
func InInt64(i64 int64, i64s []int64) bool {
	for k := range i64s {
		if i64 == i64s[k] {
			return true
		}
	}
	return false
}

// InFloat32 检查float32值在一个float32 slice中是否存在
func InFloat32(f32 float32, f32s []float32) bool {
	for k := range f32s {
		if f32 == f32s[k] {
			return true
		}
	}
	return false
}

// InFloat64 检查float64值在一个float64 slice中是否存在
func InFloat64(f64 float64, f64s []float64) bool {
	for k := range f64s {
		if f64 == f64s[k] {
			return true
		}
	}
	return false
}

// EqualBytes 检查两个[]byte是否相等
func EqualBytes(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	if (b1 == nil) != (b2 == nil) {
		return false
	}

	b2 = b2[:len(b1)]
	for i, v := range b1 {
		if v != b2[i] {
			return false
		}
	}
	return true
}
