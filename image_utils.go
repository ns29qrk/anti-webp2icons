// ==============================================================================
// File: image_utils.go
// Author: ns29qrk
// Domain: 29q.ryukyu
// Created: 2025-12-18
// License: MIT
//
// Description:
//
//	Utilities for loading (PNG/WebP) and resizing images using high-quality
//	interpolation.
//
// ==============================================================================
package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
)

// loadImage reads a PNG or WebP image from the given path.
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// image.Decode will now handle both PNG and WebP due to side-effect imports
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return img, nil
}

// resizeImage resizes the given image to the target size using high-quality Catmull-Rom interpolation.
func resizeImage(src image.Image, size int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}
