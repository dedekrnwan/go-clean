package orm

import (
	"context"
	"github.com/dedekrnwan/go-clean/internal/repository"
	"github.com/dedekrnwan/go-clean/model"
	"gorm.io/gorm"
)

type (
	User interface {
		repository.Orm[model.User, model.User]
		CountByEmail(ctx context.Context, email string) (int64, error)
	}

	user struct {
		repository.Orm[model.User, model.User]
	}
)

func NewUser(connection *gorm.DB) User {
	orm := repository.NewOrm(connection, model.User{}, model.User{})
	return &user{
		orm,
	}
}

func (m *user) CountByEmail(ctx context.Context, email string) (count int64, err error) {
	err = m.GetDBConnector().Model(model.User{}).WithContext(ctx).Where("email", email).Count(&count).Error
	return
}
