package repo

import (
	"sideq/internal/entity"
	"sideq/pkg/db"
)

type ExampleRepo struct {
	*db.Gorm
}

func New(db *db.Gorm) *ExampleRepo {
	return &ExampleRepo{db}
}

func (r *ExampleRepo) GetExamples() ([]entity.Example, error) {
	examples := []entity.Example{}
	result := r.DB.Find(&examples)
	if result.Error != nil {
		return nil, result.Error
	}
	return examples, nil
}

func (r *ExampleRepo) PostExample(example entity.Example) (entity.Example, error) {
	result := r.DB.Create(&example)
	if result.Error != nil {
		return entity.Example{}, result.Error
	}
	return example, nil
}
