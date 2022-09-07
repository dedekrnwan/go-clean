package repository

import (
	"github.com/dedekrnwan/go-clean/internal/config"
	"github.com/dedekrnwan/go-clean/internal/driven/database"
	"github.com/dedekrnwan/go-clean/internal/repository/orm"
	"gorm.io/gorm"
)

type Contract struct {
	connection *gorm.DB
	Orm        *Orm
}

func New() *Contract {
	c := &Contract{}

	///connection database init
	database.Init()

	//getting spesific db
	conn, err := database.Connection[gorm.DB](config.Config.Connection.Primary)
	if err != nil {
		panic(err)
	}

	c.connection = conn

	//init repo orm
	c.Orm = &Orm{}
	c.Orm.User = orm.NewUser(c.connection)

	return c
}
