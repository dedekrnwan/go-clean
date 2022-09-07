package orm

import (
	"context"
	"github.com/dedekrnwan/go-clean/model"
	"gorm.io/gorm"
)

type (
	User interface {
		Orm[model.User, model.User]
		CountByEmail(ctx context.Context, email string) (int64, error)
	}

	user struct {
		Orm[model.User, model.User]
	}
)

func NewUser(connection *gorm.DB) User {
	orm := NewOrm(connection, model.User{}, model.User{})
	return &user{
		orm,
	}
}

func (m *user) CountByEmail(ctx context.Context, email string) (count int64, err error) {
	//err = m.GetDBConnector().Model(model.User{})
	return
}
