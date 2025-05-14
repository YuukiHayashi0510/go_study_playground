## 比較対象

ハッシュマップをやめたuniqueSortInt（>0)でないと動作しないものと
よくある実装。

## 背景
思いつきの関数（自然数のソート関数）、秘伝のタレuniqueSortNaturalIntsをベンチマークしてみたい。

## 結果

特定の条件下でのみ、ハッシュマップを使わない処理にできるため、自然数の場合だけとてもパフォーマンスが良い。
二つの処理の差は配列orマップなので、マップは意外と遅いことがわかった。
是非とも個人用のutilsパッケージに追加したい。

### 実行コマンド

（例）
```sh
go test -bench=. -benchmem -benchtime=10000x
```

### 出力

```sh
goos: darwin
goarch: arm64
pkg: github.com/YuukiHayashi0510/go_study_playground/unique_sort_int_array_stdlib_vs_natural_ints
cpu: Apple M1 Pro
BenchmarkMyUniqueSort-10           10000              6607 ns/op           33400 B/op         13 allocs/op
BenchmarkStdUniqueSort-10          10000            104027 ns/op           82457 B/op         21 allocs/op
PASS
ok      github.com/YuukiHayashi0510/go_study_playground/unique_sort_int_array_stdlib_vs_natural_ints    1.368s
```
