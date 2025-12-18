[English](./README.md) | [日本語](./README_ja.md) | 繁體中文

# anti-webp2icons

一個簡單的 CLI 工具，可以從單張 1024x1024 的 WebP 或 PNG 圖片生成 macOS (`.icns`) 和 Windows (`.ico`) 圖標。

本工具專為 Antigravity 工作流開發，無需依賴 `iconutil` 等外部工具即可輕鬆使用。

## 特色

- **無外部依賴**: 可以在任何操作系統上為其他系統生成圖標。
- **高畫質**: 使用 Catmull-Rom 插值算法，確保圖標縮放後依然清晰。
- **跨平台**: 使用 Go 語言編寫，支持 Mac、Windows 和 Linux。
- **多種格式**: 同時生成 `.icns` (Mac)、 `.ico` (Windows) 以及多種尺寸的 PNG (Linux: 16, 32, 48, 128, 256, 512)。

## 導入方法（安裝）

請根據您的環境選擇以下其中一種方法。

### A. Go 語言使用者・工程師
在終端機執行以下指令即可完成安裝，並可在任何地方使用 `webp2icons` 指令。
```bash
go install github.com/ns29qrk/anti-webp2icons@latest
```

### B. 一般使用者・非工程師（推薦）
1. 從 [Releases](https://github.com/ns29qrk/anti-webp2icons/releases) 頁面下載適合您作業系統的檔案（Mac 請選 `macos`，Windows 請選 `windows.exe`）。
2. 將下載的檔案放在方便找的地方（例如：文件夾，或 Mac 的 `/usr/local/bin` 等）。
3. **Antigravity 使用者**: 直接將下載的檔案拖入 Antigravity，並告訴它「我想用這個製作圖標」，AI 就會幫您完成後續的設定。

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

## 與 Antigravity 整合

此存儲庫已針對 [Antigravity](https://antigravity.google) 進行優化。無需記住複雜指令，只需透過對話即可製作圖標。

1. **準備**: 按照上述「導入方法」下載或安裝工具。
2. **使用**: 在聊天框中輸入 `/create-app-icons` 並指定原始圖片即可。

> [!TIP]
> 即使不懂「設定路徑 (PATH)」等專業術語也沒關係。只要對 Antigravity 說「幫我安裝這個工具」，AI 就會檢查您的環境並協助完成設定。

## 更新歷史

### v1.0.0 (2025-12-18)
- 初版發布。
- 支援 WebP/PNG 輸入。
- 支援輸出 ICNS (macOS), ICO (Windows) 以及 Linux 用 PNG (16, 32, 48, 128, 256, 512)。
- 新增輸出控制參數 (`-mac`, `-win`, `-linux`)。
- 新增輸出目錄參數 (`-d`)。

## 免責聲明

本軟體按「原樣」提供，不提供任何形式的保證。作者不對因使用本工具而產生的任何損害承擔責任。請自行承擔使用風險。

## 授權協議

MIT 授權協議。詳情請參閱 [LICENSE](./LICENSE) 文件。
