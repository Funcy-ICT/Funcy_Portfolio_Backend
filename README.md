# Funcy-portfolio
# お約束

## プログラムを書く上で
基本的に下記のサイトに書いてることに従います。  
https://golang.org/doc/effective_go.html

## Github
#### Branch命名規則
- master
    - プロダクトとしてリリースするためのブランチ. 基本触らない
- develop(default)
    - 開発ブランチ． コードが安定し,リリース準備ができたら master へマージする. リリース前はこのブランチが最新バージョンとなる.
- feature
    - 機能の追加. develop から分岐し, develop にマージする.
    - feature-{任意で詳細}
- fix
    - 現在のプロダクトのバージョンに対する変更・修正用.
    - fix-{任意で詳細}
#### コミットメッセージ
- add:新機能
- fix:バグ修正
- wip:作業中（WIP：Work In Progress）
- clean:整理（削除も含む）

#### issue,Pull Requestのラベル(主に使って欲しいものを明記)
- bug バグの内容、解決したいことについて記述
- documentation ドキュメントの更新
- enhancement 新機能の開発
- help wanted 助けて欲しいこと(基本わからないことがあったらこれ書いて)
- question 質問、議論(わからないことではなく「これであっているのか不安だな」ということについて書いてください)
## レビュー体制
大崎、伊藤がレビュー担当


## APIドキュメント
```
make up
```
上記コマンドでコンテナ立ち上げ後にswagger UI にアクセス
[swagger UI](http://localhost:8002/)

## 検証用
web、モバイル班の方たち用になります。
これで検証用サーバを利用できます。
M1, M2 macを使用している人はdocker-compose.yml 21行目のコメントアウトを解除してください。
```
make up
# 別のターミナルを用意
make migrate-demo
```

## 実行
サーバ班
```
make up
```
### マイグレーション
```
make migrate
```
