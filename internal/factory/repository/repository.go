package repository

import (
	"github.com/dedekrnwan/go-clean/config"
	"github.com/dedekrnwan/go-clean/internal/driven/database"
	"github.com/dedekrnwan/go-clean/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	connection *gorm.DB
	User       repository.User
}

func NewFactory() *Factory {
	f := &Factory{}

	///connection database init
	database.Init()

	//getting spesific db
	conn, err := database.Connection[gorm.DB](config.Config.Connection.Primary)
	if err != nil {
		panic(err)
	}

	f.connection = conn

	//init repo
	f.User = repository.NewUser(f.connection)

	return f
}
