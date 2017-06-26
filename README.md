# GoTwitter
GoでTwitterを楽しむ

## TwitterAPIの取得
[Twitter Developer](https://dev.twitter.com/)
ここに登録をして
- Consumer_Key
- ConsumerSecret_Key
- AccessToken_Key
- AccessTokenSecret_Key
の4つをメモしておく

## TwitterAPIライブラリanacondaを利用する

```
$ go get github.com/ChimeraCoder/anaconda
```

## Config.jsonを作る
```
$ vim config.json
```
```
{
	"consumer_key": "{Your consumer key}",
	"consumer_secret": "{Your consumer secret key}",
	"access_token": "{Your token key}",
	"access_token_secret": "{Your token secret key}"
}
```

## 動かす
```
$ go run userTimeline.go -conf=config.json
```
