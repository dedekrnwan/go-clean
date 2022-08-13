package model

import "time"

type (
	Filter struct {
		Field    string      `json:"field" bson:"field"`
		Operator string      `json:"operator" bson:"operator"`
		Value    interface{} `json:"value" bson:"value"`
	}

	Pagination struct {
		Page  int `query:"page" json:"page" bson:"page"`
		Limit int `query:"limit" json:"limit" bson:"limit"`
	}

	PaginationInfo struct {
		Pagination
		Count     int64 `json:"count" bson:"count"`
		TotalPage int64 `json:"total_page" bson:"total_page"`
	}

	ResponseSingle[T any] struct {
		Data T `json:"data" bson:"data"`
	}

	ResponseMany[T any] struct {
		Data           []T            `json:"data" bson:"data"`
		PaginationInfo PaginationInfo `json:"pagination_info" bson:"pagination_info"`
	}

	Model struct {
		ID int `json:"id" gorm:"primaryKey;autoIncrement;" param:"id" swaggerignore:"true"`

		CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
		CreatedBy *int      `json:"created_by" swaggerignore:"true"`

		ModifiedAt time.Time `json:"modified_at" swaggerignore:"true"`
		ModifiedBy *int      `json:"modified_by" swaggerignore:"true"`

		DeletedAt time.Time `json:"-" gorm:"index" swaggerignore:"true"`
		DeletedBy *int      `json:"deleted_by" swaggerignore:"true"`
	}
)
