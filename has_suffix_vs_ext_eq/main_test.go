package main

import (
	"path/filepath"
	"strings"
	"testing"
)

var testCases = []struct {
	filename string
	ext      string
}{
	{"test.txt", ".txt"},
	{"document.pdf", ".pdf"},
	{"image.jpg", ".jpg"},
	{"archive.tar.gz", ".gz"},
	{"noextension", ""},
	{"hidden.file.txt", ".txt"},
	{"very.long.file.name.with.multiple.dots.pdf", ".pdf"},
	{".hiddenfile", ""},
}

func BenchmarkFilepathExt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			if filepath.Ext(tc.filename) == tc.ext {
				continue
			}
		}
	}
}

func BenchmarkStringsHasSuffix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			if strings.HasSuffix(tc.filename, tc.ext) {
				continue
			}
		}
	}
}
