# holo-back

## 前提
- go 1.14
- docker 19.3
- docker-copose 1.25
- yarn 1.22

## 概要

| name | host           | description                        |
| ---- | -------------- | ---------------------------------- |
| app  | localhost:8080 | golang, web server                 |
| db   | localhost:5432 | postgres, user=holo, pass=password |

## セットアップ

```bash
yarn setup # dockerイメージのビルド
```

## 使い方
起動
```bash
yarn start  # コンテナの起動
yarn status # コンテナのステータス確認
curl localhost:8080
```

DBマイグレーション
```bash
yarn migrate
```

停止
```bash
yarn stop   # コンテナの停止
```

## リリース用ビルド

```bash
yarn build
docker image ls
```


## 便利なエイリアス
### docker-compose
```bash
yarn docker      # docker-compose の alias
yarn docker logs # ログをみる例
```

### goose
```bash
yarn goose         # goose の alias
yarn docker status # DBの状態をみる
```
