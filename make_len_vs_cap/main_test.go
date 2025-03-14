package main

import "testing"

const (
	allocates = 1000
)

func BenchmarkMakeLen(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		arr := make([]int, allocates)
		arr[0] = 10
		_ = arr // prevent unused variable warning
	}
}

func BenchmarkMakeCap(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		arr := make([]int, 0, allocates)
		arr = append(arr, 10)
		_ = arr // prevent unused variable warning
	}
}

func BenchmarkMakeLenFill(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		arr := make([]int, allocates)
		for j := range allocates {
			arr[j] = j
		}
		_ = arr
	}
}

func BenchmarkMakeCapFill(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		arr := make([]int, 0, allocates)
		for j := range allocates {
			arr = append(arr, j)
		}
		_ = arr
	}
}
