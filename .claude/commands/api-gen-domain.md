---
description: Domain層のコード生成（Entity, ValueObject, Repository Interface）
allowed-tools: Read, Write, Glob
argument-hint: <Resource> <fields> (例: Comment body:string:1-500,author:string:1-50)
---

# Domain層生成

新しいリソースのDomain層を生成します。

## 引数
- `$ARGUMENTS` の1番目: リソース名（PascalCase）例: `Comment`
- `$ARGUMENTS` の2番目以降: フィールド定義（カンマ区切り）
  - フォーマット: `name:type:min-max`
  - 例: `body:string:1-500,author:string:1-50`
- Resource だけが引数として与えられた場合、sqlc/schema.sql からテーブルを参照し、フィールドを判断してください

## 参照ファイル
まず以下のファイルを読んで、既存のパターンを確認してください：
- `domain/entity/post.go` - Entityのパターン
- `domain/valueobject/post_id.go` - ID値オブジェクトのパターン
- `domain/valueobject/post_title.go` - 文字列制約付き値オブジェクトのパターン
- `domain/valueobject/post_title_test.go` - 値オブジェクトテストのパターン
- `domain/repository/post_repository.go` - Repositoryインターフェースのパターン

## 生成ファイル

### 1. Entity: `domain/entity/{resource}.go`

Postのパターンに従って作成：
- struct定義（ValueObjectを使用）
- `New{Resource}()` コンストラクタ
- 各フィールドのゲッターメソッド
- `Update()` メソッド
- `Reconstruct{Resource}()` 復元関数

### 2. ValueObject (ID): `domain/valueobject/{resource}_id.go`

```go
package valueobject

type {Resource}ID struct {
	value int
}

func New{Resource}ID(id int) {Resource}ID {
	return {Resource}ID{value: id}
}

func (id {Resource}ID) Int() int {
	return id.value
}
```

### 3. ValueObject (各フィールド): `domain/valueobject/{resource}_{field}.go`

制約付きフィールド（min-max指定あり）の場合：
```go
package valueobject

import "errors"

type {Resource}{Field} struct {
	value string
}

func New{Resource}{Field}({field} string) ({Resource}{Field}, error) {
	if len({field}) < {min} {
		return {Resource}{Field}{}, errors.New("{field} must be at least {min} character(s)")
	}
	if len({field}) > {max} {
		return {Resource}{Field}{}, errors.New("{field} must be at most {max} characters")
	}
	return {Resource}{Field}{value: {field}}, nil
}

func Reconstruct{Resource}{Field}({field} string) {Resource}{Field} {
	return {Resource}{Field}{value: {field}}
}

func (f {Resource}{Field}) String() string {
	return f.value
}
```

### 4. ValueObject Test: `domain/valueobject/{resource}_{field}_test.go`

制約付きフィールドのテスト：
```go
package valueobject

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew{Resource}{Field}(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid {field}",
			input:   "valid value",
			wantErr: false,
		},
		{
			name:    "empty {field}",
			input:   "",
			wantErr: true,
		},
		{
			name:    "{field} too long",
			input:   strings.Repeat("a", {max}+1),
			wantErr: true,
		},
		{
			name:    "{field} at max length",
			input:   strings.Repeat("a", {max}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := New{Resource}{Field}(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
```

### 5. Repository Interface: `domain/repository/{resource}_repository.go`

```go
package repository

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type {Resource}Repository interface {
	Save(ctx context.Context, {resource} *entity.{Resource}) (valueobject.{Resource}ID, error)
	FindByID(ctx context.Context, id valueobject.{Resource}ID) (*entity.{Resource}, error)
	FindAll(ctx context.Context) ([]*entity.{Resource}, error)
	Update(ctx context.Context, {resource} *entity.{Resource}) error
	Delete(ctx context.Context, id valueobject.{Resource}ID) error
}
```

## 生成後の確認

```bash
# テスト実行
go test ./domain/...
```

## 生成後の手動作業

生成が完了したら、以下のコマンドを実行してコードを生成してください：

```bash
# mockの作成
make mock
```
