package factory

import (
	repositoryFactory "github.com/dedekrnwan/go-clean/internal/factory/repository"
	usecaseFactory "github.com/dedekrnwan/go-clean/internal/factory/usecase"
)

type Factory struct {
	Usecase    *usecaseFactory.Factory
	Repository *repositoryFactory.Factory
}

func NewFactory() *Factory {
	f := &Factory{}
	f.Repository = repositoryFactory.NewFactory()
	f.Usecase = usecaseFactory.NewFactory(f.Repository)

	return f
}
