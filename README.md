# Gong CLI

Go言語で作成されたCLIアプリケーションです。

## インストール

```bash
go build -o gong main.go
```

## 使用方法

```bash
# ヘルプを表示
./gong help

# バージョンを表示
./gong version
```

## プロジェクト構造

```
.
├── main.go           # エントリーポイント
├── cmd/              # サブコマンド用のディレクトリ
├── internal/         # 内部パッケージ（外部からアクセス不可）
├── pkg/              # 外部からも使用可能なパッケージ
├── go.mod            # Go modules設定
└── README.md         # このファイル
```

## 開発

### ビルド

```bash
go build -o gong main.go
```

### 実行

```bash
go run main.go help
go run main.go version
```

### テスト

```bash
go test ./...
```
