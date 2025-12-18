// ==============================================================================
// File: ico.go
// Author: ns29qrk
// Domain: 29q.ryukyu
// Created: 2025-12-18
// License: MIT
//
// Description:
//
//	Logic for encoding images into Windows ICO format without external tools.
//
// ==============================================================================
package main

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"io"
)

// icoHeader represents the ICONDIR structure.
type icoHeader struct {
	Reserved uint16
	Type     uint16
	Count    uint16
}

// icoEntry represents the ICONDIRENTRY structure.
type icoEntry struct {
	Width    uint8
	Height   uint8
	ColorCnt uint8
	Reserved uint8
	Planes   uint16
	BitCount uint16
	Size     uint32
	Offset   uint32
}

// saveICO encodes images into Windows ICO format.
func saveICO(w io.Writer, images []image.Image) error {
	header := icoHeader{
		Reserved: 0,
		Type:     1, // 1 for ICON
		Count:    uint16(len(images)),
	}

	if err := binary.Write(w, binary.LittleEndian, header); err != nil {
		return err
	}

	pngBuffers := make([][]byte, len(images))
	entries := make([]icoEntry, len(images))

	offset := uint32(6 + 16*len(images)) // Header(6) + EntrySize(16)*Count

	for i, img := range images {
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return err
		}
		pngData := buf.Bytes()
		pngBuffers[i] = pngData

		bounds := img.Bounds()
		width := uint8(bounds.Dx())
		if bounds.Dx() >= 256 {
			width = 0
		}
		height := uint8(bounds.Dy())
		if bounds.Dy() >= 256 {
			height = 0
		}

		entries[i] = icoEntry{
			Width:    width,
			Height:   height,
			ColorCnt: 0,
			Reserved: 0,
			Planes:   1,
			BitCount: 32,
			Size:     uint32(len(pngData)),
			Offset:   offset,
		}
		offset += entries[i].Size
	}

	for _, entry := range entries {
		if err := binary.Write(w, binary.LittleEndian, entry); err != nil {
			return err
		}
	}

	for _, data := range pngBuffers {
		if _, err := w.Write(data); err != nil {
			return err
		}
	}

	return nil
}
