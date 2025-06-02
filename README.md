# 📔 blogAPI 📔

シンプルなブログアプリのREST API。記事の投稿や取得、コメント、いいね機能を備え、GoとMySQLで構築されています。

## ブログAPI

- 🪶 機能

  - 記事を投稿する
  - 投稿一覧を取得する
  - 特定の投稿を取得する
  - 投稿にいいねをつける
  - 投稿にコメントする

## 技術スタック

- ⚙️ アプリケーション

  - Go (API実装)
  - MySQL (RDBMS)

- 🧑‍💻 開発ツール

  - Docker (デプロイ)
  - Docker Compose (開発環境)
  - goose (DBマイグレーション)
  - (未)Github Actions (CI/CD)

## 起動方法

- 手順

  - 1..envを記述する。
  - 2.起動/停止コマンド: `make run`/`make shutdown`
  - 3.(初回起動時)Migrationの実行: `make migrate-up`

- DBのみ起動: `make db-up`
- blogAPIのみビルド: `make build`

## ディレクトリについて

- `api/`: ルーティングを記述。今後ミドルウェアの実装などもこの中に書く。
- `controllers/`: ３層アーキテクチャのプレゼンテーション層に当たる部分。ハンドラやJSONスキーマを記述してある。
- `services/`: ３層アーキテクチャのアプリケーション層に当たる部分。ブログアプリのロジックを記述してある。
- `repository/`: ３層アーキテクチャのデータ層。リポジトリパターンに近い実装を行なっている。
- `models/`: 各層で使用する共通言語として、記事とコメントを定義してある。

## エンドポイント

- `POST /article`: 記事をポストする
- `GET /article/list?page=`: 記事一覧の表示
- `GET /article/{id}`: 指定した記事の詳細を取得
- `PATCH /article/{id}`: 指定した記事にいいねする
- `POST /comment`: 記事へのコメントをポストする

