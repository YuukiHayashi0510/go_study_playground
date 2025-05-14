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
goarch: amd64
pkg: github.com/YuukiHayashi0510/go_study_playground/unique_sort_int_array_stdlib_vs_natural_ints
cpu: VirtualApple @ 2.50GHz
BenchmarkMyUniqueSort-8            10000              9211 ns/op           41592 B/op         14 allocs/op
BenchmarkStdUniqueSort-8           10000            110280 ns/op           82456 B/op         21 allocs/op
PASS
ok      github.com/YuukiHayashi0510/go_study_playground/unique_sort_int_array_stdlib_vs_natural_ints    1.822s
```
