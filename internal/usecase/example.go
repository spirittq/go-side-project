package usecase

import (
	"context"
	"sideq/internal/entity"
)

type ExampleUseCase struct {
	repo ExampleRepo
}

func New(r ExampleRepo) *ExampleUseCase {
	return &ExampleUseCase{
		repo: r,
	}
}

func (uc *ExampleUseCase) GetExamples(ctx context.Context) ([]entity.Example, error) {
	examples, err := uc.repo.GetExamples()
	if err != nil {
		return nil, err
	}
	return examples, nil
}

func (uc *ExampleUseCase) PostExample(ctx context.Context, examplePostRequest entity.ExamplePostRequest) (entity.Example, error) {
	example := entity.Example{
		Field1: examplePostRequest.Field1,
		Field2: examplePostRequest.Field2,
		Field3: examplePostRequest.Field3,

	}
	example, err := uc.repo.PostExample(example)
	if err != nil {
		return entity.Example{}, err
	}
	return example, nil
}
