package gfnv

import (
	"hash/fnv"
	"testing"
)

func BenchmarkStandardFnv32(b *testing.B) {
	msg := []byte("hello world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f := fnv.New32()
		f.Write(msg)
		_ = f.Sum32()
	}
}

func BenchmarkGfnv32(b *testing.B) {
	msg := []byte("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fnv32(msg)
	}
}

func BenchmarkStandardFnv64(b *testing.B) {
	msg := []byte("hello world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f := fnv.New64()
		f.Write(msg)
		_ = f.Sum64()
	}
}

func BenchmarkGfnv64(b *testing.B) {
	msg := []byte("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fnv64(msg)
	}
}
