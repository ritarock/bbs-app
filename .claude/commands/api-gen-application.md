---
description: Application層のコード生成（DTO, Usecase + テスト）
allowed-tools: Read, Write, Glob
argument-hint: <Resource> <fields> (例: Comment body:string,author:string)
---

# Application層生成

新しいリソースのApplication層（DTO、Usecase、テスト）を生成します。

## 引数
- `$ARGUMENTS` の1番目: リソース名（PascalCase）例: `Comment`
- `$ARGUMENTS` の2番目以降: フィールド定義（カンマ区切り）
  - フォーマット: `name:type` （制約はApplication層では不要）
  - 例: `body:string,author:string`
- Resource だけが引数として与えられた場合、domain/entity, domain/repository, domain/valueobject を参照し、判断してください

## 前提条件
- Domain層が生成済みであること
- Mock（testing/mock/{resource}_repository.go）が生成済みであること

## 参照ファイル
まず以下のファイルを読んで、既存のパターンを確認してください：
- `application/dto/post_dto.go` - DTOのパターン
- `application/usecase/post/create_post.go` - Usecaseのパターン
- `application/usecase/post/create_post_test.go` - Usecaseテストのパターン
- `application/usecase/post/get_post.go`
- `application/usecase/post/list_posts.go`
- `application/usecase/post/update_post.go`
- `application/usecase/post/delete_post.go`

## 生成ファイル

### 1. DTO: `application/dto/{resource}_dto.go`

```go
package dto

import "time"

type Create{Resource}Input struct {
	{Field1} string
	{Field2} string
	// ... 各フィールド
}

type Create{Resource}Output struct {
	ID        int
	{Field1}  string
	{Field2}  string
	CreatedAt time.Time
}

type Get{Resource}Input struct {
	ID int
}

type Get{Resource}Output struct {
	ID        int
	{Field1}  string
	{Field2}  string
	CreatedAt time.Time
}

type {Resource}Item struct {
	ID        int
	{Field1}  string
	{Field2}  string
	CreatedAt time.Time
}

type List{Resource}Output struct {
	{Resources} []{Resource}Item
}

type Update{Resource}Input struct {
	ID       int
	{Field1} string
	{Field2} string
}

type Update{Resource}Output struct {
	ID        int
	{Field1}  string
	{Field2}  string
	CreatedAt time.Time
}

type Delete{Resource}Input struct {
	ID int
}
```

### 2. Usecase: `application/usecase/{resource}/create_{resource}.go`

```go
package {resource}

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
)

type Create{Resource}Usecase struct {
	{resource}Repo repository.{Resource}Repository
}

func NewCreate{Resource}Usecase({resource}Repo repository.{Resource}Repository) *Create{Resource}Usecase {
	return &Create{Resource}Usecase{
		{resource}Repo: {resource}Repo,
	}
}

func (u *Create{Resource}Usecase) Execute(ctx context.Context, input dto.Create{Resource}Input) (*dto.Create{Resource}Output, error) {
	{resource}, err := entity.New{Resource}(input.{Field1}, input.{Field2})
	if err != nil {
		return nil, err
	}

	{resource}ID, err := u.{resource}Repo.Save(ctx, {resource})
	if err != nil {
		return nil, err
	}

	return &dto.Create{Resource}Output{
		ID:        {resource}ID.Int(),
		{Field1}:  {resource}.{Field1}().String(),
		{Field2}:  {resource}.{Field2}().String(),
		CreatedAt: {resource}.CreatedAt(),
	}, nil
}
```

### 3. 他のUsecase

同様のパターンで以下を作成：
- `application/usecase/{resource}/get_{resource}.go` - FindByIDで取得
- `application/usecase/{resource}/list_{resources}.go` - FindAllで取得
- `application/usecase/{resource}/update_{resource}.go` - FindByID → Update → 保存
- `application/usecase/{resource}/delete_{resource}.go` - Deleteを呼び出し

### 4. Usecase Test: `application/usecase/{resource}/create_{resource}_test.go`

```go
package {resource}_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/{resource}"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate{Resource}Usecase_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    dto.Create{Resource}Input
		mockFunc func(m *mock.Mock{Resource}Repository)
		want     *dto.Create{Resource}Output
		hasError bool
	}{
		// テストケース
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockRepo := mock.NewMock{Resource}Repository(ctrl)
			test.mockFunc(mockRepo)
			uc := {resource}.NewCreate{Resource}Usecase(mockRepo)
			got, err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// アサーション
			}
		})
	}
}
```

### 5. 他のUsecaseテスト

同様のパターンで以下を作成：
- `application/usecase/{resource}/get_{resource}_test.go`
- `application/usecase/{resource}/list_{resources}_test.go`
- `application/usecase/{resource}/update_{resource}_test.go`
- `application/usecase/{resource}/delete_{resource}_test.go`

## 生成後の確認

```bash
# テスト実行
go test ./application/...
```
