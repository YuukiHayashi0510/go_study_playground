package main

import (
	"sort"
	"testing"
)

// arrayだけで実装(バリデーションでsort.Ints使用)
func uniqueSortNaturalInts(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	// 配列をコピーしてソート
	result := make([]int, len(arr))
	copy(result, arr)
	sort.Ints(result)

	if result[0] < 0 {
		return result
	}

	// 最大値を見つける
	maxVal := 0
	for _, v := range result {
		if v > maxVal {
			maxVal = v
		}
	}

	// 最大値+1の長さのスライスを作成
	sorted := make([]int, maxVal+1)
	for _, v := range result {
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
