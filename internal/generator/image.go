package generator

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"

	"github.com/QiZD90/resume-generator/internal/entity"
)

func ImageToBase64(src []byte) (string, error) {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	_, format, err := image.DecodeConfig(bytes.NewReader(src))
	if err != nil {
		return "", err
	}

	base64.StdEncoding.Encode(dst, src)

	return fmt.Sprintf("data:image/%s;base64, %s", format, dst), nil
}

func CropAndEncodeImage(src []byte, crop entity.Crop) (string, error) {
	return ImageToBase64(src)
}
