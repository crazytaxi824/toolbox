package test

import (
	"convert"
	"testing"
)

func TestStrToByte(t *testing.T) {
	str := "fjdskjkhjhkjhkjfsdlkjf"
	b1 := convert.StrToByte(str)
	b2 := []byte(str)
	l := len(str)
	for i := 0; i < l; i++ {
		if b1[i] != b2[i] {
			t.Error("failed")
		}
	}
}

func BenchmarkStrToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert.StrToByte("s")
	}
}

func BenchmarkByteToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert.BytesToStr([]byte{64})
	}
}
