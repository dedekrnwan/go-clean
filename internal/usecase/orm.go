package usecase

import (
	"context"
	"github.com/dedekrnwan/go-clean/internal/repository"
	"github.com/dedekrnwan/go-clean/model"
)

type (
	Orm[T any, Y any] interface {
		Count(ctx context.Context, filters []model.Filter) (int64, error)
		Find(ctx context.Context, search string, filters []model.Filter, ascending []string, descending []string, pagination model.Pagination, preloads []string, excludesOrder ...string) ([]Y, *model.PaginationInfo, error)
		FindOne(ctx context.Context, id int, preloads []string) (*Y, error)
		CreateOne(ctx context.Context, data *Y) (*Y, error)
		CreateMany(ctx context.Context, data []Y) ([]Y, error)
		UpdateOne(ctx context.Context, id int, data *Y) (*Y, error)
		DeleteOne(ctx context.Context, id int) error
	}

	orm[T any, Y any] struct {
		model repository.Orm[T, Y]
	}
)

func NewOrm[T any, Y any](model repository.Orm[T, Y]) Orm[T, Y] {
	return &orm[T, Y]{
		model,
	}
}

func (u *orm[T, Y]) Count(ctx context.Context, filters []model.Filter) (int64, error) {
	return u.model.Count(ctx, filters)
}

func (u *orm[T, Y]) Find(ctx context.Context, search string, filters []model.Filter, ascending []string, descending []string, pagination model.Pagination, preloads []string, excludesOrder ...string) ([]Y, *model.PaginationInfo, error) {
	return u.model.Find(ctx, search, filters, ascending, descending, pagination, preloads, excludesOrder...)
}

func (u *orm[T, Y]) FindOne(ctx context.Context, id int, preloads []string) (*Y, error) {
	return u.model.FindOne(ctx, id, preloads)
}

func (u *orm[T, Y]) CreateOne(ctx context.Context, data *Y) (*Y, error) {
	return u.model.CreateOne(ctx, data)
}

func (u *orm[T, Y]) CreateMany(ctx context.Context, data []Y) ([]Y, error) {
	return u.model.CreateMany(ctx, data)
}

func (u *orm[T, Y]) UpdateOne(ctx context.Context, id int, data *Y) (*Y, error) {
	return u.model.UpdateOne(ctx, id, data)
}

func (u *orm[T, Y]) DeleteOne(ctx context.Context, id int) error {
	return u.model.DeleteOne(ctx, id)
}
