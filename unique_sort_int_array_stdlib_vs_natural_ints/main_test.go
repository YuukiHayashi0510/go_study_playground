package main

import (
	"sort"
	"testing"
)

// arrayだけで実装(バリデーションでsort.Ints使用)
func uniqueSortNaturalInts(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	sort.Ints(arr)
	if arr[0] < 0 {
		return arr
	}

	sorted := make([]int, len(arr)+1)
	for _, v := range arr {
		sorted[v] = v
	}

	uniqueList := make([]int, 0)
	for k, v := range sorted {
		if v > 0 {
			uniqueList = append(uniqueList, k)
		}
	}

	return uniqueList
}

// mapとsort.Intsを使った方
func stdUniqueSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	uniqueMap := make(map[int]bool)
	for _, num := range arr {
		uniqueMap[num] = true
	}

	uniqueList := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		uniqueList = append(uniqueList, k)
	}

	sort.Ints(uniqueList)

	return uniqueList
}

func BenchmarkMyUniqueSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1000)
		for j := range arr {
			arr[j] = j
		}
		uniqueSortNaturalInts(arr)
	}
}

func BenchmarkStdUniqueSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1000)
		for j := range arr {
			arr[j] = j
		}
		stdUniqueSort(arr)
	}
}
