# WEB-AI-SPEACKER
例えば主婦の方ならキッチンやリビングに物理的なAIスピーカーを置くことにはメリットがあります。
しかし我々エンジニア、どうせ常時パソコンの前にいるのですからブラウザでAIスピーカーが使えればもう物理的なAIスピーカーを置く必要も無いんじゃないでしょうか。

これはそんな方に贈る、誰でも簡単にブラウザ上で使えるAIスピーカーです。

![image](https://user-images.githubusercontent.com/25472671/65835454-37bbef00-e321-11e9-85a4-a3670d18eb6f.png)

## How to use
```
    go run main.go
    npm run serve

  cp web/dist/.index.html web/dist/index.html
  cp web/public/.index.html web/public/index.html
```

# TODO
## 優先度 高
- 常駐化
  [x] シンプルなｊｓのイベント常駐
  [x] トグルスイッチ
- [ ] herokuデプロイ

## 優先度 中
- [x] コンポーネント分割しっかりする
- [x] アクセスキー等の隠蔽
- [ ] Firebaseによる履歴機能、wakeupワードの設定

## 優先度 低
- [ ] コトバンクや、流行語のための検索、トレンドワードに強いAPIでの検索も入れたい
- [ ] 多機能化(天気やニュース、タイマーなど)
    キメラサービスにならないように注意
- [ ] 検索した言葉に合った画像を持ってくる

## 優先度 未定
- [ ] 常時起動で使用制限がかかった場合はsnowボーイの導入を行う

## Memo
CLI上でのログイン
heroku login --interactive

npm installはプロジェクトルートで実行されるためにpackageを移す

## ワイとHeroku、〜デプロイへの軌跡〜
- herokuへのpush
```
$ sudo heroku container:push --recursive
```
- herokuへのデプロイ
```
$ heroku container:release web
```

- heroku上でのdocker状況確認
```
$ heroku ps
```

- スケール変更
```
$ heroku ps:scale web=1
$ heroku ps:scale backend=1
```

- クラッシュしていて止めたい時
```
$ heroku restart
$ heroku ps:scale web=0
$ heroku ps:scale backend=0
```

- LocalでProcfileからサービスを起動する方法
  Procfileが存在しない場合はherokuが自動的に言語を検出してくれる
```
$ heroku local web
```

- ビルドが長引く問題とか
```
$ heroku builds:cancel
```

- herokuのログ
```
$ heroku logs -t 
```

- 明日の自分へ
デプロイで吐かれるginのエラーを解決せよ
```
2019-10-09T12:23:54.84613+00:00 app[web.1]: invalid character 'h' looking for beginning of value
2019-10-09T12:23:54.846132+00:00 app[web.1]: /api/controller/speech_recog.go:94 (0x992abd)
2019-10-09T12:23:54.846133+00:00 app[web.1]: SpeechPost.func1: panic(err)
```

- ビルドタイプの変更(nodejs指定だとpackage.jsonを要求されたりする)
```
$ heroku buildpacks:clear
```

- デプロイタイプの変更
```
$ heroku container:push web
```
