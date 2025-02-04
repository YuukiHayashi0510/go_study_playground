## 比較対象

Golang のプリミティブ型である Map の探索処理
`for _, v := range` vs map["hoge"]

## 背景

Go における Map の探索について、パフォーマンスが気になったため。

## 結果

map["hoge"]の探索の方が速い

### 実行コマンド

```
go test -bench=. -benchmem -benchtime=10000x
```

### 出力

```
goos: darwin
goarch: arm64
pkg: github.com/YuukiHayashi0510/go_study_playground/map_range_vs_lookup
cpu: Apple M1 Pro
BenchmarkMapRange-10               10000             54773 ns/op           32372 B/op        320 allocs/op
BenchmarkMapLookup-10              10000               140.7 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/YuukiHayashi0510/go_study_playground/map_range_vs_lookup     0.804s
```
