English | [日本語](./README_ja.md) | [繁體中文](./README_zh-TW.md)

# anti-webp2icons

Simple CLI tool to generate macOS (`.icns`) and Windows (`.ico`) icons from a single 1024x1024 WebP or PNG image.

Developed to be used within Antigravity workflows without external dependencies like `iconutil`.

## Features

- **No External Dependencies**: Run on any OS to generate icons for any other OS.
- **High Quality**: Uses Catmull-Rom interpolation for sharp icon resizing.
- **Cross-Platform**: Built with Go, supports Mac, Windows, and Linux.
- **Multiple Formats**: Generates `.icns` (Mac), `.ico` (Windows), and multiple sizes of PNGs (Linux: 16, 32, 48, 128, 256, 512).

## Installation

### From Source
```bash
go install github.com/p_suke/anti-webp2icons@latest
```

### From Releases
Download the pre-built binary for your platform from the [Releases](https://github.com/p_suke/anti-webp2icons/releases) page.

## Usage

### Basic Usage
webp2icons my_icon.webp

# Specify output directory and base name
webp2icons -i source.png -d ./assets -o app_icon
```

### Output Control
You can specify which formats to generate using flags.
```bash
# Generate only macOS icon
webp2icons -mac my_icon.webp

# Generate macOS and Windows icons
webp2icons -mac -win my_icon.webp
```

### Options
- `-i`: Input file path.
- `-o`: Output base name (without extension).
- `-d`: Output directory path.
- `-mac`: Generate macOS .icns
- `-win`: Generate Windows .ico
- `-linux`: Generate Linux .png
- `-v`: Show version.
- `-h`: Show help.

## License

MIT License. See [LICENSE(optional)](#) for details.
