// ==============================================================================
// File: icns.go
// Author: ns29qrk
// Domain: 29q.ryukyu
// Created: 2025-12-18
// License: MIT
//
// Description:
//
//	Logic for encoding images into macOS ICNS format using Apple OSType markers.
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

// icnsType maps size to its ICNS type identifier.
var icnsType = map[int]string{
	16:   "is32", // also can use 'icp4'
	32:   "il32", // also can use 'icp5'
	64:   "icp6",
	128:  "it32",
	256:  "ic08",
	512:  "ic09",
	1024: "ic10",
}

// saveICNS encodes images into macOS ICNS format.
func saveICNS(w io.Writer, images map[int]image.Image) error {
	var body bytes.Buffer

	// Support Retina (represented as pairs like 16 + 32, but for modern macOS, PNG icons are preferred)
	// We'll focus on the standard PNG-based ICNS entries.
	for size, img := range images {
		ostype, ok := icnsType[size]
		if !ok {
			continue
		}

		var pngBuf bytes.Buffer
		if err := png.Encode(&pngBuf, img); err != nil {
			return err
		}
		data := pngBuf.Bytes()

		// Write OSType
		_ = binary.Write(&body, binary.BigEndian, [4]byte{ostype[0], ostype[1], ostype[2], ostype[3]})
		// Write Length (8 bytes header + data)
		_ = binary.Write(&body, binary.BigEndian, uint32(len(data)+8))
		// Write Data
		_, _ = body.Write(data)
	}

	// Write Global Header
	// magic: 'icns' (4 bytes)
	_, _ = w.Write([]byte("icns"))
	// total size: 8 bytes header + body size
	_ = binary.Write(w, binary.BigEndian, uint32(body.Len()+8))
	// Write Body
	_, _ = body.WriteTo(w)

	return nil
}
