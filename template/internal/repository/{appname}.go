package repository

import (
	"{appname}/internal/model"
)

type DemoRepository interface {
	FirstById(id int64) (*model.Demo, error)
}
type demoRepository struct {
	*Repository
}

func NewDemoRepository(repository *Repository) DemoRepository {
	return &demoRepository{
		Repository: repository,
	}
}

func (r *demoRepository) FirstById(id int64) (*model.Demo, error) {
	var demo model.Demo
	// TODO: query db
	return &demo, nil
}
