package qrgen

import (
	"bytes"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func Createimage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}

func GenQRCode(url string, width, height int) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, width, height)
}
