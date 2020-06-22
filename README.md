# holo-back

## 前提
- docker 19.3
- docker-copose 1.25
- yarn 1.22
- heroku 7.39

## 概要

| name | host           | description                        |
| ---- | -------------- | ---------------------------------- |
| app  | localhost:8080 | golang, web server                 |
| db   | localhost:5432 | postgres, user=holo, pass=password |

## セットアップ

```bash
yarn install
```

以下のファイルにAPIキーを設定する
```
./scripts/app/secret.env
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

ログの確認
```bash
yarn log -f
```

停止
```bash
yarn stop   # コンテナの停止
```

テスト
```bash
yarn lint
yarn test
```

クリーン
```bash
yarn clean
```

## 便利なエイリアス
### docker-compose
```bash
yarn docker      # docker-compose の alias
yarn docker logs # ログをみる例
```

### goose
```bash
yarn migrate        # goose の alias
yarn migrate status # DBの状態をみる
```

### publish js-api
```bash
cd scripts/js-api
yarn build
yarn publish --access public
```
