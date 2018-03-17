package xis

// InStr 检查string值在一个string slice中是否存在
func InStr(str string, s []string) bool {
	for k := range s {
		if str == s[k] {
			return true
		}
	}
	return false
}

// InInt 检查int值在一个int slice中是否存在
func InInt(is int, i []int) bool {
	for k := range i {
		if is == i[k] {
			return true
		}
	}
	return false
}

// InInt64 检查int64值在一个int64 slice中是否存在
func InInt64(is int64, i []int64) bool {
	for k := range i {
		if is == i[k] {
			return true
		}
	}
	return false
}

// InFloat32 检查float32值在一个float32 slice中是否存在
func InFloat32(fs float32, f []float32) bool {
	for k := range f {
		if fs == f[k] {
			return true
		}
	}
	return false
}

// InFloat64 检查float64值在一个float64 slice中是否存在
func InFloat64(fs float64, f []float64) bool {
	for k := range f {
		if fs == f[k] {
			return true
		}
	}
	return false
}

// EqualBytes 检查两个[]byte是否相等
func EqualBytes(b1, b2 []byte) bool {
	l1 := len(b1)
	l2 := len(b2)
	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
