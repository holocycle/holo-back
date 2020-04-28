# holo-back

## 前提
- go 1.14
- docker 19.3
- docker-copose 1.25

## 使い方 (Local)

依存ライブラリのダウンロード
```bash
go mod download
```

ローカルビルド & 実行
```bash
go run main.go
curl localhost:8080
```

## 使い方 (Docker)

ビルド
```bash
docker-compose build
```

起動
```bash
docker-compose up -d
curl localhost:8080
```
