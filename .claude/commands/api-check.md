---
description: BBS APIの動作確認を行う
allowed-tools: Bash(curl:*)
argument-hint: [endpoint] (posts | posts/:id | create | update)
---

# BBS API 動作確認

以下のAPIエンドポイントを確認してください。

## 利用可能なエンドポイント

### 投稿一覧取得 (posts)
```bash
curl -s http://localhost:8080/posts | jq .
```

### 投稿詳細取得 (posts/:id)
```bash
curl -s http://localhost:8080/posts/{id} | jq .
```

### 投稿作成 (create)
```bash
curl -s -X POST http://localhost:8080/posts \
   -H "Content-Type: application/json" \
   -d '{"title":"テスト投稿","content":"これはテスト内容です"}' | jq .
```

### 投稿更新 (update)
```bash
curl -s -X PUT http://localhost:8080/posts/{id} \
   -H "Content-Type: application/json" \
   -d '{"title":"更新後タイトル","content":"更新後の内容"}' | jq .
```

## タスク

引数 `$ARGUMENTS` に基づいて、適切なAPIエンドポイントを実行してください。

- `posts` または引数なし: 投稿一覧を取得
- `posts/1` や `1` など数値を含む: 該当IDの投稿詳細を取得
- `create`: テスト投稿を作成
- `update 1` など: 該当IDの投稿を更新

結果を整形して表示し、レスポンスの内容を簡潔に説明してください。
