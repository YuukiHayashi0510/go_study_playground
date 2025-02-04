## 比較対象

Golang 標準パッケージの strings.HasSuffix と path/filepath.Ext

## 背景

Golang 標準パッケージの strings.HasSuffix と filepath.Ext のどちらが速いか気になったから

## 結果

`strings.HasSuffix`の方が速い

### 実行コマンド

```
go test -bench=. -benchmem -benchtime=10000x
```

### 出力

```
goos: darwin
goarch: arm64
pkg: github.com/YuukiHayashi0510/go_study_playground/has_suffix_vs_ext_eq
cpu: Apple M1 Pro
BenchmarkFilepathExt-10                    10000               103.6 ns/op             0 B/op          0 allocs/op
BenchmarkStringsHasSuffix-10               10000                62.91 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/YuukiHayashi0510/go_study_playground/has_suffix_vs_ext_eq    0.305s
```
