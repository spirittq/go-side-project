package usecase

import (
	"context"
	"sideq/internal/entity"
)

type Example interface {
	GetExamples(ctx context.Context) ([]entity.Example, error)
	PostExample(ctx context.Context, example entity.ExamplePostRequest) (entity.Example, error)
}

type ExampleRepo interface {
	GetExamples() ([]entity.Example, error)
	PostExample(entity.Example) (entity.Example, error)
}
