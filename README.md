# WEB-AI-SPEACKER
[![Netlify Status](https://api.netlify.com/api/v1/badges/93d960d4-1bf1-496e-b74b-dd2cda8ce949/deploy-status)](https://app.netlify.com/sites/web-ai-speaker/deploys)


例えば主婦の方ならキッチンやリビングに物理的なAIスピーカーを置くことにはメリットがあります。  
しかし我々エンジニア、どうせ常時パソコンの前にいるのですからブラウザでAIスピーカーが使えればもう物理的なAIスピーカーを置く必要も無いんじゃないでしょうか。  

これはそんな方に贈る、誰でも簡単にブラウザ上で使えるAIスピーカーです。  

![image](https://user-images.githubusercontent.com/25472671/65835454-37bbef00-e321-11e9-85a4-a3670d18eb6f.png)

## How to use
please see Makefile

# TODO
## 優先度 高
- 常駐化
  - [x] シンプルなjsのイベント常駐
  - [x] トグルスイッチ
- [x] herokuデプロイ
- [x] netlifyデプロイ
- [x] サービス名変更
    ゲーム中に横のディスプレイに表示するっていうようなタイトルにする
    game, search, ai, Display side assistant

## 優先度 中 (デプロイしてカッコがつくぐらい)
- [x] コンポーネント分割しっかりする
- [x] アクセスキー等の隠蔽

- [x] favアイコン変更
- [x] タイトル変更

- [x] デプロイした時に文字サイズやGridSystemが変わることの原因調査
    他のPCのローカル環境でやってみるとデプロイした方との差異は無かった
    恐らく自分のパソコンで良くない設定をしてしまっているんじゃないかと思われる。
- [x] デプロイした時にタブに出るサイト名がVueAppになるのをWebAiSpeakerにする
- [x] スライダーボタンをロードアイコン付きボタンに変更
- [x] 音声認識結果がとりあえず表示される機能
- [ ] DocomoAPIでの読み上げ
    どこまで読み上げるかのロジックが発生してしまうので注意
    また、クレデンシャルデータを扱うためにpushできない
- [x] WakeupWordをとりあえず短くする
    利便性の話
- [x] BGMを音声で流す
    フルスクリーンでゲームしてるとブラウザ操作するのが面倒だから
- [x] googleの音声合成に対応させる
    音声合成on off のボタン欲しくない？
    - [x] サイドバーとりあえず実装
        上部にずらずらボタンがあるのは嫌なのでスマートにサイドバー
- [x] 音声認識ロジック分離、vueイベントバスを利用
    ttsとsttを分離する！
- [x] WakeupWordを短くするか無くすかでどっちがいいかをヒアリングしてみる
    なくして誤爆があるかどうかを調べて見る必要があるかも
    あったほうがよさげ
- [x] WakeupWord追加・削除のロジック実装
- [x] サイドバーをかっちょよくする  
    この時デザインについて調べてから実装に移ると吉
    - [x] 読み上げon,off
    - [x] wakeupワードのlist、削除ボタン付き
- [ ] コトバンク対応で一気に語彙力を増やす
- [x] BGMを音声で止めたり音量操作したりする
- [ ] youtube再生されないバグを直す
- [ ] アイコン

## 優先度 低 (より便利)
- [x] 音声認識ロジック分離、*.js
    vue->外部jsの呼び出しは可能だったが、外部jsからライブラリ(axios)を呼び出すとwebpack周りのエラーが出る。
    webpackの設定をいじればあるいはエラーは解消できそうだがwebpack周りはvue-cliが管理してくれていて、そこを触るのはリスキーだと考えて棚上げする.
- [x] 検索中の言葉が画面に表示されるとかloadアイコンが回るとかの実装
- [x] 長過ぎる音声入力テキストは画面に表示しないようにする
    監視されている感が強くなってしまうため
- [x] WakeupWord追加ボタン
    ボタン押した時にテキストエリアをクリアする
- [x] 使い方のシーンを書いてあるページ作成?
    冒頭に仕様シーンが出ていて、下にスクロールして使い出す感じ
- [ ] Firebase
    履歴機能、wakeupワードの設定
- [ ] 流行語のための検索、トレンドワードに強いAPIでの検索
- [ ] 検索した言葉に合った画像を持ってくる
- [ ] VueAppってところどころで出るのを直す
    分からん。聞こう。
- [ ] デプロイするとアイコンが表示されなくなる問題を解消
- [ ] アイコン作る
- [ ] クッキーかlocalstrageを使って登録ワードの記憶
- [ ] ハロー あいみょん を流して
- [ ] タイトルバー周辺の装飾

## 優先度 未定
- [ ] 常時起動で使用制限がかかった場合はsnowボーイの導入を行う
- [ ] 多機能化?(天気やニュース、タイマーなど)
    キメラサービスにならないように注意
- [ ] クレデンシャル情報
    一般公開可能なデータにしかアクセスできないものとはいえAPIキーを晒している
    使われても金銭的ダメージは発生しないが、リクエスト上限もあるので隠した方が吉
    
## Memo
CLI上でのログイン
heroku login --interactive

npm installはプロジェクトルートで実行されるためにpackageを移す

```
  cp web/dist/.index.html web/dist/index.html
  cp web/public/.index.html web/public/index.html
```

## Herokuコマンド例
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
ローカルでは通るがデプロイ時では吐かれるginのエラーを解決せよ
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
