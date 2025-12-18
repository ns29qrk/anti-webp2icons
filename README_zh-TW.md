[English](./README.md) | [日本語](./README_ja.md) | 繁體中文

# anti-webp2icons

一個簡單的 CLI 工具，可以從單張 1024x1024 的 WebP 或 PNG 圖片生成 macOS (`.icns`) 和 Windows (`.ico`) 圖標。

本工具專為 Antigravity 工作流開發，無需依賴 `iconutil` 等外部工具即可輕鬆使用。

## 特色

- **無外部依賴**: 可以在任何操作系統上為其他系統生成圖標。
- **高畫質**: 使用 Catmull-Rom 插值算法，確保圖標縮放後依然清晰。
- **跨平台**: 使用 Go 語言編寫，支持 Mac、Windows 和 Linux。
- **多種格式**: 同時生成 `.icns` (Mac)、 `.ico` (Windows) 以及多種尺寸的 PNG (Linux: 16, 32, 48, 128, 256, 512)。

## 安裝方法

### 從源碼編譯
```bash
go install github.com/p_suke/anti-webp2icons@latest
```

### 下載執行檔
請從 [Releases](https://github.com/p_suke/anti-webp2icons/releases) 頁面下載適用於您平台的預編譯二進制文件。

## 使用方法

### 基本用法
工具將使用輸入文件名作為基礎，自動生成各種格式。
```bash
webp2icons my_icon.webp
```

### 進階用法 (指定輸出目錄與檔案名稱)
使用 `-d` 指定目錄， `-o` 指定檔案基礎名稱。
```bash
webp2icons -i source.png -d ./assets -o app_icon
```

### 輸出格式控制
您可以使用參數指定要生成的格式。
```bash
# 僅生成 macOS 圖標
webp2icons -mac my_icon.webp

# 同時生成 macOS 和 Windows 圖標
webp2icons -mac -win my_icon.webp
```

### 參數說明
- `-i`: 輸入文件路徑。
- `-o`: 輸出基礎名稱（不含擴展名）。
- `-d`: 輸出目錄路徑。
- `-mac`: 生成 macOS .icns
- `-win`: 生成 Windows .ico
- `-linux`: 生成 Linux .png
- `-v`: 顯示版本號。
- `-h`: 顯示說明。

## 授權協議

MIT 授權協議。詳情請參閱代碼文件頭部。
