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

## 実行
```
make run
```
### マイグレーション
```
make migrate
```
golang-migrate環境がない方はこちら
```
//make run後
$ docker container exec -it funcy_portfolio_backend-api-1 bash

$ make migrate-demo
```
成功していれば以下のような表示
```
20221026122655/u create_users (20.390913ms)
20221105092401/u work (60.712116ms)
20230103140307/u alter_users_status (89.638866ms)
```