package test

import (
	"testing"
	"xis"
)

func TestInStrSuccess(t *testing.T) {
	str := "abc"
	s := []string{"jfkdls", "fjdsl", "ookogt", "tryuetwuy", "abc"}
	if !xis.InStr(str, s) {
		t.Error("failed")
	}
}

func TestInStrFailed(t *testing.T) {
	str := "abcd"
	s := []string{"jfkdls", "fjdsl", "ookogt", "tryuetwuy", "abc"}
	if xis.InStr(str, s) {
		t.Error("failed")
	}
}

func BenchmarkInStr(b *testing.B) {
	str := "abc"
	s := []string{"jfkdls", "fjdsl", "ookogt", "tryuetwuy", "abc"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xis.InStr(str, s)
	}
}

func TestInIntSuccess(t *testing.T) {
	i := 12
	s := []int{21, 43, 76, 87, 05, 12}
	if !xis.InInt(i, s) {
		t.Error("failed")
	}
}

func TestInIntFailed(t *testing.T) {
	i := 123
	s := []int{21, 43, 76, 87, 05, 12}
	if xis.InInt(i, s) {
		t.Error("failed")
	}
}

func BenchmarkInInt(b *testing.B) {
	is := 12
	s := []int{21, 43, 76, 87, 05, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xis.InInt(is, s)
	}
}

func TestBEBSuccess(t *testing.T) {
	a := []byte("abc")
	b := []byte("abc")
	if !xis.EqualBytes(a, b) {
		t.Error("failed")
	}
}

func TestBEBFailed(t *testing.T) {
	a := []byte("abcf")
	b := []byte("abcd")
	if xis.EqualBytes(a, b) {
		t.Error("failed")
	}
}

func BenchmarkBEB(b *testing.B) {
	a := []byte("abcf")
	c := []byte("abc")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xis.EqualBytes(a, c)
	}
}
