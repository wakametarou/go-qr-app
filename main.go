package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"go-qr-app/qrgen"
)

func main() {
	outputPath := flag.String("o", "./image.png", "Path to output file")
	width := flag.Int("w", 200, "Width of the output image of the QR Code (px)")
	depth := flag.Int("d", 200, "height of the output image of the QR Code (px)")

	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		fmt.Println("URL is empty")
		return
	}

	file, err := os.Create(*outputPath)
	if err != nil {
		fmt.Printf("file generation failed: %v\n", err)
		return
	}
	defer file.Close()

	img, err := qrgen.GenQRCode(url, *width, *depth)
	if err != nil {
		fmt.Printf("image generation failed: %v\n", err)
		return
	}

	png.Encode(file, img)
}
