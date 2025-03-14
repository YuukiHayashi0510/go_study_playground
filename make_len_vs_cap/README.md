## 比較対象

## 背景
makeを使った、lenとcapのパフォーマンスが気になったため。

## 結果

### 実行コマンド

```
go test -bench=. -benchmem -benchtime=10000x
```

### 出力
```
goos: darwin
goarch: arm64
pkg: github.com/YuukiHayashi0510/go_study_playground/make_len_vs_cap
cpu: Apple M1 Pro
BenchmarkMakeLen-10                10000                 0.4292 ns/op          0 B/op          0 allocs/op
BenchmarkMakeCap-10                10000                 0.4250 ns/op          0 B/op          0 allocs/op
BenchmarkMakeLenFill-10            10000               420.7 ns/op             0 B/op          0 allocs/op
BenchmarkMakeCapFill-10            10000               804.3 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/YuukiHayashi0510/go_study_playground/make_len_vs_cap 0.390s
```
