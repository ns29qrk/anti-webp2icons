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

### A. For Go Users (Recommended)
If you have Go installed, you can install the tool globally with:
```bash
go install github.com/ns29qrk/anti-webp2icons@latest
```

### B. For Non-Go Users
1. Download the pre-built binary for your platform from the [Releases](https://github.com/ns29qrk/anti-webp2icons/releases) page.
2. Place the binary in a folder of your choice (e.g., `/usr/local/bin` on Mac).
3. **Antigravity Users**: You can simply drag and drop the downloaded binary into Antigravity and ask it to "Use this to create icons," and the AI will help you with the setup.

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

## Antigravity Integration

This repository is optimized for [Antigravity](https://antigravity.google). You can automate icon creation by simply chatting.

1. **Setup**: Download or install the tool as described in the "Installation" section.
2. **Usage**: Use the `/create-app-icons` command in your chat and specify your source image.

> [!TIP]
> If you are not familiar with technical terms like "PATH," just ask Antigravity to "Help me install this tool," and the AI will support your setup.

## Changelog

### v1.0.0 (2025-12-18)
- Initial release.
- Supports WebP/PNG input.
- Outputs ICNS (macOS), ICO (Windows), and PNG (Linux: 16, 32, 48, 128, 256, 512).
- Added output control flags (`-mac`, `-win`, `-linux`).
- Added output directory flag (`-d`).

## Disclaimer

This software is provided "as is," without warranty of any kind. The author shall not be liable for any damages arising from the use of this tool. Use it at your own risk.

## License

MIT License. See the [LICENSE](./LICENSE) file for details.
