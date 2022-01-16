package embedded

import (
	"embed"
	"encoding/base64"
	"fmt"
	"net/http"
)

var (
	//go:embed static/*
	staticFS embed.FS
)

type (
	Image struct {
		Source         string
		Alt            string
		base64Encoding string
		HTML           string
	}
)

var (
	errImg = Image{
		HTML: "<img src=\"\" alt=\"Error loading image\" />",
	}
)

func NewImg(source, alt string) Image {
	b, err := staticFS.ReadFile(source)
	if err != nil {
		return errImg
	}

	var b64 string
	mimeType := http.DetectContentType(b)
	switch mimeType {
	case "image/jpeg":
		b64 = "data:image/jpeg;base64,"
	case "image/png":
		b64 = "data:image/png;base64,"
	}
	b64 += base64.StdEncoding.EncodeToString(b)

	return Image{
		Source:         source,
		Alt:            alt,
		base64Encoding: b64,
		HTML:           fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />", b64, alt),
	}
}
