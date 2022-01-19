package embedded

import (
	"embed"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	ImgPathGopher      = "static/gopher.png"
	ImgPathRDS         = "static/Res_Amazon-Aurora_Amazon-RDS-Instance_48_Dark.png"
	ImgPathCloudWatch  = "static/Res_Amazon-CloudWatch_Alarm_48_Dark.png"
	ImgPathDynamoDB    = "static/Res_Amazon-DynamoDB_Table_48_Dark.png"
	ImgPathEC2         = "static/Res_Amazon-EC2_Instances_48_Dark.png"
	ImgPathElastiCache = "static/Res_Amazon-ElastiCache_ElastiCache-for-Redis_48_Dark.png"
	ImgPathS3          = "static/Res_Amazon-Simple-Storage-Service_S3-Standard_48_Dark.png"
	ImgPathMyFace      = "static/my_face.png"
	ImgPathYTLogo      = "static/youtube_logo.png"
	ImgPathTwitterLogo = "static/twitter_logo.png"
)

var (
	//go:embed static/*
	staticFS embed.FS
)

type (
	Image struct {
		Source         string
		Alt            string
		Base64Encoding string
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
		Base64Encoding: b64,
		HTML:           fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />", b64, alt),
	}
}
