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
### 機能に関する部分
基本的に機能などに関するPRはサーバー班の未来大メンバー1人、同じサービスの他のプラットフォームのメンバー1人の計２名に確認して問題がなくなった時点でマージ
### 技術に関係ないドキュメント関係
同じサービスの他のサーバー班の確認が得られた時点でマージ

## 実行するために
### データベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

- Macの場合
```cassandraql
$ export PORT=8080 
    MYSQL_USER=root
    MYSQL_PASSWORD=admin 
    MYSQL_HOST=127.0.0.1 
    MYSQL_PORT=3306 
    MYSQL_DATABASE=funcy
```
- Windowsの場合
```cassandraql
$ SET PORT=8080
$ SET MYSQL_USER=root
$ SET MYSQL_PASSWORD=admin
$ SET MYSQL_HOST=192.168.99.100
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=funcy
```
>docker-compose up -d   
go run  main.go