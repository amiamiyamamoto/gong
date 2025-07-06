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
