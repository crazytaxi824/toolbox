package test

import (
	"refactor"
	"testing"
	"xis"
)

func TestMerge(t *testing.T) {
	r, err := refactor.MergeBytesBak([]byte("a"), []byte("b"), []byte("c"))
	if err != nil {
		t.Error(err)
		return
	}
	if !xis.EqualBytes(r, []byte("abc")) {
		t.Error("failed")
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		refactor.MergeBytesBak([]byte("Abf"), []byte("Abf"), []byte("Abf"), []byte("Abf"), []byte("Abf"))
	}
}

func BenchmarkMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		refactor.MergeBytes([]byte("Abf"), []byte("Abf"), []byte("Abf"), []byte("Abf"), []byte("Abf"))
	}
}
