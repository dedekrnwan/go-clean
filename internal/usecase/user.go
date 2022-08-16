package usecase

import (
	repositoryFactory "github.com/dedekrnwan/go-clean/internal/factory/repository"
	repositoryOrm "github.com/dedekrnwan/go-clean/internal/repository"
	"github.com/dedekrnwan/go-clean/model"
)

type (
	User interface {
		Orm[model.User, model.User]
	}

	user struct {
		Orm[model.User, model.User]

		repositoryUser repositoryOrm.User
	}
)

func NewUser(
	factory *repositoryFactory.Factory,
) User {
	return &user{
		Orm:            NewOrm[model.User, model.User](factory.User),
		repositoryUser: factory.User,
	}
}
