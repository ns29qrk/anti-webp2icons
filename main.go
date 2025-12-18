// ==============================================================================
// File: main.go
// Author: ns29qrk
// Domain: 29q.ryukyu
// Created: 2025-12-18
// License: MIT
//
// Description:
//
//	CLI entry point for anti-webp2icons project. Orchestrates the conversion
//	process from a single source image to multiple app icon formats.
//
// ==============================================================================
package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	var (
		inputPath  string
		outputPath string
		outDir     string
		showVer    bool
		mac        bool
		win        bool
		linux      bool
	)

	flag.StringVar(&inputPath, "i", "", "Input WebP or PNG file path")
	flag.StringVar(&outputPath, "o", "", "Output base name (without extension)")
	flag.StringVar(&outDir, "d", "", "Output directory path")
	flag.BoolVar(&showVer, "v", false, "Show version")
	flag.BoolVar(&mac, "mac", false, "Generate macOS .icns")
	flag.BoolVar(&win, "win", false, "Generate Windows .ico")
	flag.BoolVar(&linux, "linux", false, "Generate Linux .png")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: webp2icons [options] <input_file>\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if showVer {
		fmt.Printf("webp2icons version %s\n", version)
		os.Exit(0)
	}

	// Helper to parse flags from positional arguments if users put them in the wrong place
	args := flag.Args()
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "-i" && i+1 < len(args):
			inputPath = args[i+1]
			i++
		case arg == "-o" && i+1 < len(args):
			outputPath = args[i+1]
			i++
		case arg == "-d" && i+1 < len(args):
			outDir = args[i+1]
			i++
		case arg == "-mac":
			mac = true
		case arg == "-win":
			win = true
		case arg == "-linux":
			linux = true
		case !strings.HasPrefix(arg, "-") && inputPath == "":
			inputPath = arg
		}
	}

	if inputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if outDir != "" {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}

	// If no specific format is selected, default to all
	if !mac && !win && !linux {
		mac, win, linux = true, true, true
	}

	if err := run(inputPath, outputPath, outDir, mac, win, linux); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully.")
}

func run(inputPath, outBase, outDir string, mac, win, linux bool) error {
	// 1. Load source image
	src, err := loadImage(inputPath)
	if err != nil {
		return err
	}

	// Determine output base name
	if outBase == "" {
		outBase = strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	}

	// 2. Prepare images for each format
	icoSizes := []int{16, 32, 48, 256}
	icnsSizes := []int{16, 32, 64, 128, 256, 512, 1024}
	linuxSizes := []int{16, 32, 48, 128, 256, 512}

	// Collect unique sizes to minimize resize operations
	sizeMap := make(map[int]image.Image)
	var allSizes []int
	if win {
		allSizes = append(allSizes, icoSizes...)
	}
	if mac {
		allSizes = append(allSizes, icnsSizes...)
	}
	if linux {
		allSizes = append(allSizes, linuxSizes...)
	}

	for _, s := range allSizes {
		if _, ok := sizeMap[s]; !ok {
			sizeMap[s] = resizeImage(src, s)
		}
	}

	// 3. Save ICO
	if win {
		destPath := filepath.Join(outDir, outBase+".ico")
		icoFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer icoFile.Close()

		var icoImages []image.Image
		for _, s := range icoSizes {
			icoImages = append(icoImages, sizeMap[s])
		}
		if err := saveICO(icoFile, icoImages); err != nil {
			return err
		}
		fmt.Printf("Generated: %s\n", destPath)
	}

	// 4. Save ICNS
	if mac {
		destPath := filepath.Join(outDir, outBase+".icns")
		icnsFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer icnsFile.Close()

		icnsImages := make(map[int]image.Image)
		for _, s := range icnsSizes {
			icnsImages[s] = sizeMap[s]
		}
		if err := saveICNS(icnsFile, icnsImages); err != nil {
			return err
		}
		fmt.Printf("Generated: %s\n", destPath)
	}

	// 5. Save Linux PNGs
	if linux {
		for _, s := range linuxSizes {
			pngName := fmt.Sprintf("%s_icon_%d.png", outBase, s)
			destPath := filepath.Join(outDir, pngName)
			f, err := os.Create(destPath)
			if err != nil {
				return err
			}
			if err := png.Encode(f, sizeMap[s]); err != nil {
				f.Close()
				return err
			}
			f.Close()
			fmt.Printf("Generated: %s\n", destPath)
		}
	}

	return nil
}
