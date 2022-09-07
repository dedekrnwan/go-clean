package usecase

import (
	repositoryContract "github.com/dedekrnwan/go-clean/internal/contract/repository"
	"github.com/dedekrnwan/go-clean/internal/usecase"
)

type Contract struct {
	User usecase.User
}

func New(c *repositoryContract.Contract) *Contract {
	u := &Contract{}
	u.User = usecase.NewUser(c)
	return u
}
