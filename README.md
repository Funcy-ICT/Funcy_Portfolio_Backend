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
未定

## メールの確認方法について
以下の[URL](http://localhost:8025/)にアクセスしてメールの受信を確認してください
```
http://localhost:8025/
```


## APIドキュメント
### 設定が必要な環境変数
(`*`: 必須項目)
| * | NAME             | Default | Description            |
| - | ---------------- | ------- | ---------------------- |
|   | SUPER_ACCOUNT_ID |         | スーパアカウント用のID |


### スーパアカウント
`SUPER_ACCOUNT_ID`を指定すると、指定したユーザとしてログインできる実質的に有効期限のないTokenが発行されます。
Tokenは、標準出力に、
```
[00] yyyy/mm/dd hh:nn:ss SuperAccountID: {SUPER_ACCOUNT_ID}
[00] yyyy/mm/dd hh:nn:ss SuperAccountToken: {Access Token}
```
という形で出力されます。

### 立ち上げ
```
make up
```
上記コマンドでコンテナ立ち上げ後にswagger UI にアクセス
[swagger UI](http://localhost:8002/)

## 検証用
web、モバイル班の方たち用になります。
dev/1ブランチを使用します。
dev/1に移動後に以下を実行。
これで検証用サーバを利用できます。
```
make up
# 別のターミナルを用意
make maigrate-demo
```

## 実行
```
make run
```
### マイグレーション
```
make migrate
```