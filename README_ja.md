[English](./README.md) | 日本語 | [繁體中文](./README_zh-TW.md)

# anti-webp2icons

1枚の 1024x1024 WebP または PNG 画像から、macOS (`.icns`) および Windows (`.ico`) 用のアイコンを生成するシンプルな CLI ツールです。

`iconutil` などの外部ツールに依存せず、Antigravity のワークフローなどで手軽に利用できるように開発されました。

## 特徴

- **外部依存なし**: macOS以外のOS（WindowsやLinux）でも Mac 向けアイコンを生成できます。
- **高品質**: Catmull-Rom 補間を使用して、美しくシャープにリサイズします。
- **クロスプラットフォーム**: Go 製のため、Mac、Windows、Linux で動作します。
- **マルチフォーマット**: `.icns` (Mac)、 `.ico` (Windows)、および Linux 用の各サイズ PNG (16, 32, 48, 128, 256, 512) を生成します。
- **AI フレンドリー**: Antigravity 用のワークフローを内蔵しています。

## 導入方法（インストール）

あなたの環境に合わせて、以下の2つの方法から選んでください。

### A. Go 言語を使っている・エンジニアの方
以下のコマンドをターミナルで実行するだけでインストールが完了し、どこからでも `webp2icons` コマンドが使えるようになります。
```bash
go install github.com/ns29qrk/anti-webp2icons@latest
```

### B. 一般ユーザー・非エンジニアの方（推奨）
1. [Releases](https://github.com/ns29qrk/anti-webp2icons/releases) ページから、お使いの OS に合ったファイルをダウンロードします（Mac なら `macos`、Windows なら `windows.exe` と書かれたもの）。
2. ダウンロードしたファイルを、分かりやすい場所（例：書類フォルダや、Mac なら `/usr/local/bin` など）に置きます。
3. **Antigravity をお使いの場合**: ダウンロードしたファイルを Antigravity にドラッグ＆ドロップして、「これを使ってアイコンを作りたい」と伝えるだけで、あとの難しい設定は AI が代行してくれます。

## 使い方

### 基本的な使い方
入力ファイル名をベースに各フォーマットのアイコンが生成されます。
```bash
webp2icons my_icon.webp
```

### 詳細な指定 (出力ディレクトリとファイル名の指定)
`-d` でディレクトリ、`-o` でファイル名のベースを指定できます。
```bash
webp2icons -i source.png -d ./assets -o app_icon
```

### 出力OSの制御
特定のOS用のみ出力したい場合は、以下のフラグを指定します（複数指定可）。
```bash
# macOS用のみ出力
webp2icons -mac my_icon.webp

# macOSとWindows用を出力
webp2icons -mac -win my_icon.webp
```

### オプション
- `-i`: 入力ファイルのパス
- `-o`: 出力ファイルのベース名（拡張子なし）
- `-d`: 出力ディレクトリのパス
- `-mac`: macOS 用アイコン (.icns) を生成
- `-win`: Windows 用アイコン (.ico) を生成
- `-linux`: Linux 用アイコン (PNG) を生成
- `-v`: バージョンを表示
- `-h`: ヘルプを表示

## Antigravity との連携

このリポジトリは [Antigravity](https://antigravity.google) に最適化されています。難しいコマンドを覚えなくても、チャットで指示するだけでアイコンが作れます。

1. **準備**: 上記の「導入方法」に従ってファイルをダウンロードするかインストールします。
2. **使う**: チャット欄で `/create-app-icons` と入力し、元になる画像を指定するだけです。

> [!TIP]
> 「パスを通す」といった専門用語が分からなくても、Antigravity に「インストールを手伝って」と言えば、AI があなたの環境を調べて設定をサポートしてくれます。

## 更新履歴

### v1.0.0 (2025-12-18)
- 初版リリース。
- WebP/PNG 入力に対応。
- ICNS (macOS), ICO (Windows), および Linux 用 PNG (16, 32, 48, 128, 256, 512) の出力をサポート。
- 出力制御フラグ (`-mac`, `-win`, `-linux`) を追加。
- 出力先ディレクトリ指定フラグ (`-d`) を追加。

## 免責事項

本ソフトウェアは「現状のまま」提供され、いかなる保証もありません。作者は本ツールの使用から生じるいかなる損害についても責任を負いません。自己責任でご利用ください。

## ライセンス

MIT ライセンス。詳細は [LICENSE](./LICENSE) ファイルをご確認ください。
