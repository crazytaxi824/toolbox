package refactor

import "bytes"

// MergeBytes 多个[]byte合并成一个[]byte
func MergeBytes(baseBytes []byte, bytesSlice ...[]byte) (result []byte) {
	result = baseBytes
	for _, v := range bytesSlice {
		result = append(result, v...)
	}
	return
}

// MergeBytesBak 多个[]byte合并成一个[]byte
func MergeBytesBak(baseBytes []byte, bytesSlice ...[]byte) (result []byte, err error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	_, err = buf.Write(baseBytes)
	if err != nil {
		return
	}
	for k := range bytesSlice {
		_, err = buf.Write(bytesSlice[k])
		if err != nil {
			return
		}
	}
	result = buf.Bytes()
	return
}
