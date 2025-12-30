---
description: Infrastructure層のコード生成（Repository実装, Handler + テスト）
allowed-tools: Read, Write, Glob
argument-hint: <Resource> (例: Comment)
---

# Infrastructure層生成

新しいリソースのInfrastructure層（Repository実装、Handler）を生成します。

## 引数
- `$ARGUMENTS`: リソース名（PascalCase）例: `Comment`

## 前提条件
- Domain層が生成済みであること
- Application層が生成済みであること
- `sqlc generate` が実行済みであること（`infra/database/query/` にクエリが生成済み）
- `make tsp.compile` が実行済であること
- `make ogen` が実行済みであること（`infra/handler/api/` にAPIが生成済み）

## 参照ファイル
まず以下のファイルを読んで、既存のパターンを確認してください：
- `infra/database/post_repository.go` - Repository実装のパターン
- `infra/handler/handler.go` - 統合Handlerのパターン
- `infra/handler/post_handler.go` - 個別Handlerのパターン
- `infra/database/query/*.go` - sqlc生成のクエリ関数
- `infra/handler/api/*.go` - ogen生成のAPI型

## 生成ファイル

### 1. Repository実装: `infra/database/{resource}_repository.go`

```go
package database

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/infra/database/query"
)

type {Resource}Repository struct {
	queries *query.Queries
}

func New{Resource}Repository(db query.DBTX) repository.{Resource}Repository {
	return &{Resource}Repository{
		queries: query.New(db),
	}
}

func (r *{Resource}Repository) Save(ctx context.Context, {resource} *entity.{Resource}) (valueobject.{Resource}ID, error) {
	result, err := r.queries.Insert{Resource}(ctx, query.Insert{Resource}Params{
		{Field1}:   {resource}.{Field1}().String(),
		{Field2}:   {resource}.{Field2}().String(),
		CreatedAt: {resource}.CreatedAt(),
	})
	if err != nil {
		return valueobject.{Resource}ID{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return valueobject.{Resource}ID{}, err
	}

	return valueobject.New{Resource}ID(int(id)), nil
}

func (r *{Resource}Repository) FindByID(ctx context.Context, id valueobject.{Resource}ID) (*entity.{Resource}, error) {
	result, err := r.queries.Select{Resource}(ctx, int64(id.Int()))
	if err != nil {
		return nil, err
	}
	return r.toEntity(result), nil
}

func (r *{Resource}Repository) FindAll(ctx context.Context) ([]*entity.{Resource}, error) {
	result, err := r.queries.Select{Resources}(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.{Resource}, len(result))
	for i, row := range result {
		items[i] = r.toEntity(row)
	}

	return items, nil
}

func (r *{Resource}Repository) Update(ctx context.Context, {resource} *entity.{Resource}) error {
	return r.queries.Update{Resource}(ctx, query.Update{Resource}Params{
		ID:       int64({resource}.ID().Int()),
		{Field1}: {resource}.{Field1}().String(),
		{Field2}: {resource}.{Field2}().String(),
	})
}

func (r *{Resource}Repository) Delete(ctx context.Context, id valueobject.{Resource}ID) error {
	return r.queries.Delete{Resource}(ctx, int64(id.Int()))
}

func (r *{Resource}Repository) toEntity(row query.{Resource}) *entity.{Resource} {
	return entity.Reconstruct{Resource}(
		valueobject.New{Resource}ID(int(row.ID)),
		row.{Field1},
		row.{Field2},
		row.CreatedAt,
	)
}
```

### 2. Handler: `infra/handler/{resource}_handler.go`

```go
package handler

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/{resource}"
	"github.com/ritarock/bbs-app/infra/handler/api"
)

type {Resource}Handler struct {
	api.UnimplementedHandler
	create{Resource}Usecase *{resource}.Create{Resource}Usecase
	get{Resource}Usecase    *{resource}.Get{Resource}Usecase
	list{Resource}Usecase   *{resource}.List{Resource}Usecase
	update{Resource}Usecase *{resource}.Update{Resource}Usecase
	delete{Resource}Usecase *{resource}.Delete{Resource}Usecase
}

var _ api.Handler = (*{Resource}Handler)(nil)

func New{Resource}Handler(
	create{Resource}Usecase *{resource}.Create{Resource}Usecase,
	get{Resource}Usecase *{resource}.Get{Resource}Usecase,
	list{Resource}Usecase *{resource}.List{Resource}Usecase,
	update{Resource}Usecase *{resource}.Update{Resource}Usecase,
	delete{Resource}Usecase *{resource}.Delete{Resource}Usecase,
) *{Resource}Handler {
	return &{Resource}Handler{
		create{Resource}Usecase: create{Resource}Usecase,
		get{Resource}Usecase:    get{Resource}Usecase,
		list{Resource}Usecase:   list{Resource}Usecase,
		update{Resource}Usecase: update{Resource}Usecase,
		delete{Resource}Usecase: delete{Resource}Usecase,
	}
}

func (h *{Resource}Handler) {Resources}Create(ctx context.Context, req *api.Create{Resource}Request) (*api.{Resource}, error) {
	input := dto.Create{Resource}Input{
		{Field1}: req.Get{Field1}(),
		{Field2}: req.Get{Field2}(),
	}

	output, err := h.create{Resource}Usecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.{Resource}{
		ID:        int64(output.ID),
		{Field1}:  output.{Field1},
		{Field2}:  output.{Field2},
		CreatedAt: output.CreatedAt,
	}, nil
}

// ... 他のメソッド (Read, List, Update, Delete) も同様に実装
```

**注意**: 個別Handlerには `NewError` メソッドを実装しません。エラーハンドリングは統合Handlerで行います。

### 3. 統合Handler更新: `infra/handler/handler.go`

既存の統合Handlerに新しいリソースのHandlerを追加します。

```go
// Handler構造体に新しいリソースを追加
type Handler struct {
	*PostHandler
	*CommentHandler
	*{Resource}Handler  // 追加
}

// NewHandler関数の引数と初期化に追加
func NewHandler(
	postHandler *PostHandler,
	commentHandler *CommentHandler,
	{resource}Handler *{Resource}Handler,  // 追加
) *Handler {
	return &Handler{
		PostHandler:      postHandler,
		CommentHandler:   commentHandler,
		{Resource}Handler: {resource}Handler,  // 追加
	}
}

// NewErrorメソッドに新しいリソースのエラーケースを追加
func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	// ... 既存のswitch文に追加
	case errMsg == "{resource} not found":
		statusCode = http.StatusNotFound
		message = "{Resource} not found"
	// バリデーションエラーの条件にも追加
	case strings.Contains(errMsg, "title") || strings.Contains(errMsg, "content") || strings.Contains(errMsg, "{field}"):
		statusCode = http.StatusBadRequest
		message = errMsg
}
```

## 生成後の確認

```bash
# テスト実行
go test ./infra/...
```
