package main

import (
	"testing"
)

// b.N - количество запусков бенчмарка
func BenchmarkInserXInMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100)
	}
}

func BenchmarkInserXInMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(1000)
	}
}

func BenchmarkInserXInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100000)
	}
}

func BenchmarkInserXInMapInterface100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100)
	}
}

func BenchmarkInserXInMapInterface1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(1000)
	}
}

func BenchmarkInserXInMapInterface100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100000)
	}
}

func BenchmarkInserXInSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100)
	}
}

func BenchmarkInserXInSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(1000)
	}
}

func BenchmarkInserXInSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100000)
	}
}