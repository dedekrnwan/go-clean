package repository

import (
	repositoryFactory "github.com/dedekrnwan/go-clean/internal/factory/repository"
	"github.com/dedekrnwan/go-clean/internal/usecase"
)

type Factory struct {
	User usecase.User
}

func NewFactory(factoryRepository *repositoryFactory.Factory) *Factory {
	f := &Factory{}
	f.User = usecase.NewUser(factoryRepository)
	return f
}
