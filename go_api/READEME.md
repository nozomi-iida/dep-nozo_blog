# nozo blog のバックエンド

## db のセットアップ

1. sqlite3 のファイルを作成する

```.
$ sqlite3
> tmp/data.db;
```

2. マイグレートを実行する

```.
$ docker compose exec app bash
> sql-migrate up
```
