package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/YuukiHayashi0510/go-utils/empty"
)

const (
	templateDir    = "templates"
	readmeTemplate = "README_TEMPLATE"
	testTemplate   = "test_template"
	readmeTarget   = "README.md"
	testTarget     = "main_test.go"
	dirPermission  = 0755
)

// ベンチマークを作成するにあたって必要なものを用意するCLIツール
// 必要なもの：結果をまとめるためのREADMEと、ベンチマークを実行するためのコードのボイラープレートが入ったフォルダ
// 使い方：`go run main.go make_len_vs_cap`
// これで、make_len_vs_capフォルダが作成され、その中にREADME.mdとmain_test.goがコピーされる
func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("folder created successfully")
}

func run() error {
	// 引数のバリデーション
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: <folder_name>\nExample: make_len_vs_cap")
	}

	folderName := os.Args[1]

	// バリデーション
	if empty.Is(folderName) {
		return fmt.Errorf("folder name cannot be empty")
	}
	if _, err := os.Stat(folderName); err == nil {
		return fmt.Errorf("folder '%s' already exists", folderName)
	}

	// フォルダの作成
	if err := os.Mkdir(folderName, dirPermission); err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}

	// READMEのコピー
	tmpReadmeSrc := filepath.Join(templateDir, readmeTemplate)
	tmpReadmeDst := filepath.Join(folderName, readmeTarget)
	if err := copyFile(tmpReadmeSrc, tmpReadmeDst); err != nil {
		return err
	}

	// テストファイルのコピー
	tmpTestSrc := filepath.Join(templateDir, testTemplate)
	tmpTestDst := filepath.Join(folderName, testTarget)
	if err := copyFile(tmpTestSrc, tmpTestDst); err != nil {
		return err
	}

	return nil
}

// ファイルをコピーする関数
func copyFile(src, dst string) (retErr error) {
	// srcファイルを開く
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file '%s': %w", src, err)
	}
	defer func() {
		if closeErr := srcFile.Close(); closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("failed to close source file '%s': %w", src, closeErr)
		}
	}()

	// dstファイルを作成
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file '%s': %w", dst, err)
	}
	defer func() {
		if closeErr := dstFile.Close(); closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("failed to close destination file '%s': %w", dst, closeErr)
		}
	}()

	// srcファイルの内容をdstファイルにコピー
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy from '%s' to '%s': %w", src, dst, err)
	}

	return nil
}
