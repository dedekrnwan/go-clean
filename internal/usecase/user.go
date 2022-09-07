package usecase

import (
	repositoryContract "github.com/dedekrnwan/go-clean/internal/contract/repository"
	repositoryOrm "github.com/dedekrnwan/go-clean/internal/repository/orm"
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
	r *repositoryContract.Contract,
) User {
	return &user{
		Orm:            NewOrm[model.User, model.User](r.Orm.User),
		repositoryUser: r.Orm.User,
	}
}
