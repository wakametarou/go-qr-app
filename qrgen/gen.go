package qrgen

import (
	"bytes"
	"image"
	"image/png"
)

func Createimage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}
