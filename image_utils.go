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

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

// loadImage reads a PNG or WebP image from the given path.
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Try decoding as WebP first, then fall back to standard image.Decode (for PNG)
	img, err := webp.Decode(file)
	if err == nil {
		return img, nil
	}

	// Reset file pointer for the next attempt
	_, _ = file.Seek(0, 0)
	img, _, err = image.Decode(file)
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
