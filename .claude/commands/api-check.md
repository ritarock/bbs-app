---
description: BBS APIの動作確認を行う
allowed-tools: Bash(curl:*)
argument-hint: [endpoint] (posts | posts/:id | create | update | delete | comments | comment-create | comment-update | comment-delete | signup | signin | me)
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

### 投稿削除 (delete)
```bash
curl -s -X DELETE http://localhost:8080/posts/{id}
```

---

## コメントAPI

### コメント一覧取得 (comments)
```bash
curl -s http://localhost:8080/posts/{postId}/comments | jq .
```

### コメント詳細取得 (comments/:id)
```bash
curl -s http://localhost:8080/posts/{postId}/comments/{id} | jq .
```

### コメント作成 (comment-create)
```bash
curl -s -X POST http://localhost:8080/posts/{postId}/comments \
   -H "Content-Type: application/json" \
   -d '{"body":"これはテストコメントです"}' | jq .
```

### コメント更新 (comment-update)
```bash
curl -s -X PUT http://localhost:8080/posts/{postId}/comments/{id} \
   -H "Content-Type: application/json" \
   -d '{"body":"更新後のコメント"}' | jq .
```

### コメント削除 (comment-delete)
```bash
curl -s -X DELETE http://localhost:8080/posts/{postId}/comments/{id}
```

---

## 認証API

### ユーザー登録 (signup)
```bash
curl -s -X POST http://localhost:8080/auth/signup \
   -H "Content-Type: application/json" \
   -d '{"email":"test@example.com","password":"password123"}' | jq .
```

### サインイン (signin)
```bash
curl -s -X POST http://localhost:8080/auth/signin \
   -H "Content-Type: application/json" \
   -d '{"email":"test@example.com","password":"password123"}' | jq .
```

### 現在のユーザー情報取得 (me)
```bash
curl -s http://localhost:8080/auth/me \
   -H "Authorization: Bearer {token}" | jq .
```

## タスク

引数 `$ARGUMENTS` に基づいて、適切なAPIエンドポイントを実行してください。

### 投稿API
- `posts` または引数なし: 投稿一覧を取得
- `posts/1` や `1` など数値を含む: 該当IDの投稿詳細を取得
- `create`: テスト投稿を作成
- `update 1` など: 該当IDの投稿を更新
- `delete 1` など: 該当IDの投稿を削除

### コメントAPI
- `comments 1`: 投稿ID=1のコメント一覧を取得
- `comments 1 2`: 投稿ID=1のコメントID=2の詳細を取得
- `comment-create 1`: 投稿ID=1にテストコメントを作成
- `comment-update 1 2`: 投稿ID=1のコメントID=2を更新
- `comment-delete 1 2`: 投稿ID=1のコメントID=2を削除

### 認証API
- `signup`: テストユーザーを登録
- `signup test2@example.com password456`: 指定したメールアドレスとパスワードで登録
- `signin`: テストユーザーでサインイン
- `signin test2@example.com password456`: 指定したメールアドレスとパスワードでサインイン
- `me {token}`: トークンを使って現在のユーザー情報を取得

結果を整形して表示し、レスポンスの内容を簡潔に説明してください。
