package main

import (
	"path/filepath"
	"strings"
	"testing"
)

const (
	MimeTypeXml         = "text/xml"
	MimeTypePlainTxt    = "text/plain"
	MimeTypeCSS         = "text/css"
	MimeTypeJpeg        = "image/jpeg"
	MimeTypeGif         = "image/gif"
	MimeTypePng         = "image/png"
)

const (
	XMLExt      = ".xml"
	PlainTxtExt = ".txt"
	CSSExt      = ".css"
	JpegExt     = ".jpeg"
	gifExt      = ".gif"
	PngExt      = ".png"
)

var mimeTypeMap = map[string]string{
	CSSExt:  MimeTypeCSS,
	JpegExt: MimeTypeJpeg,
	gifExt:  MimeTypeGif,
	PngExt:  MimeTypePng,
}

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

func BenchmarkMapRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			for _, mimeType := range mimeTypeMap {
				if strings.HasSuffix(tc.filename, tc.ext) {
					b.Log(mimeType)
					continue
				}
			}
		}
	}
}

func BenchmarkMapLookup(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			if mimeType, exists := mimeTypeMap[filepath.Ext(tc.filename)]; exists {
				b.Log(mimeType)
				continue
			}
		}
	}
}
