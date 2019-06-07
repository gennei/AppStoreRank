# AppStoreRank

## What

App Store のランキングをSlackへ通知します

## Usage

```sh
$ mv .env.sample .env
```

エディタで `.env` の中身を編集してください。

```
$ cat .env.sample
APP_ID=id1233456
SLACK_USERNAME=表示名
SLACK_ICON_EMOJI=アイコン絵文字
SLACK_CHANNEL=#general
INCOMING_WEBHOOK_URL=https://hooks.slack.com/services/xxxxxxxxxxxxxxxxxxx
```

### execute
```sh
$ go run main.go
```

in slack

> 読書管理ブクログ - 本棚/読書記録 は 「ブック」内75位 です

## env

AWS lambda で実行することを想定しています。
.env ファイルを編集せずに、AWS lambda の環境変数を設定しても動きます。