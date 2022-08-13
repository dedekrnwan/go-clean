package repository

import (
	"github.com/dedekrnwan/go-clean/config"
	"github.com/dedekrnwan/go-clean/internal/driven/database"
	repositoryOrm "github.com/dedekrnwan/go-clean/internal/repository/orm"
	"gorm.io/gorm"
)

type Factory struct {
	connection *gorm.DB
	User       repositoryOrm.User
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
	f.User = repositoryOrm.NewUser(f.connection)

	return f
}
