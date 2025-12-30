---
description: API仕様生成（TypeSpec + SQL定義）
allowed-tools: Read, Write, Edit
argument-hint: <resource> <fields> (例: comment body:string:1-500,author:string:1-50)
---

# API仕様生成

新しいAPIリソースのTypeSpec定義とSQL定義を生成します。

## 引数
- `$ARGUMENTS` の1番目: リソース名（単数形、小文字）例: `comment`
- `$ARGUMENTS` の2番目以降: フィールド定義（カンマ区切り）
  - フォーマット: `name:type:min-max`
  - 例: `body:string:1-500,author:string:1-50`

## 参照ファイル
まず以下のファイルを読んで、既存のパターンを確認してください：
- `main.tsp` - TypeSpec定義のパターン
- `sqlc/schema.sql` - テーブル定義のパターン
- `sqlc/query.sql` - SQLクエリのパターン

## 生成内容

### 1. Schema (sqlc/schema.sql に追記)
- ユーザが使っているDBを確認すること(例: Mysql, PostgreSQL, SQLite)
- id は引数として渡されていなくても auto increment で作成するかユーザに確認すること

### 2. Query (sqlc/query.sql に追記)
- ユーザに必要なSQLをすること (例: CRUDだけでよいか？、READ は他に必要な検索条件がないか？)

以下を参考に、SQLを定義:

```sql
-- name: Insert{Resource} :execresult
INSERT INTO {resources} (
    {fields}
) VALUES (
    ?
);

-- name: Select{Resource} :one
SELECT * FROM {resources}
WHERE id = ?
;

-- name: Select{Resources} :many
SELECT * FROM {resources}
ORDER BY id
;

-- name: Update{Resource} :exec
UPDATE {resources}
SET {field1} = ?, {field2} = ?
WHERE id = ?
;

-- name: Delete{Resource} :exec
DELETE FROM {resources}
WHERE id = ?
```

### 3. TypeSpec (main.tsp に追記)
- 2. Query (sqlc/query.sql に追記) で生成した内容から、必要なAPIを作ること
- PATH についてはユーザに確認をすること

以下を参考に、新しいリソースの定義を追記：

```typespec
model {Resource} {
  id: int64;
  // 各フィールド（制約付き）
  @minLength(min) @maxLength(max) {field}: string;
  createdAt: utcDateTime;
}

model {Resource}List {
  items: {Resource}[]
}

model Create{Resource}Request {
  // 各フィールド（制約付き）
}

model Update{Resource}Request {
  // 各フィールド（制約付き）
}

@route("/{resources}")
interface {Resources} {
  @post create(@body body: Create{Resource}Request): {
    @statusCode statusCode: 201;
    @body body: {Resource};
  } | Error;

  @get read(@path id: int64): {Resource} | Error;
  @get list(): {Resource}List | Error;

  @put update(@path id: int64, @body body: Update{Resource}Request): {Resource} | Error;

  @delete delete(@path id: int64): {
    @statusCode statusCode: 204;
  } | Error;
}
```
