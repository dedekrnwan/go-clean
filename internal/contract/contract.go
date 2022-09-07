package contract

import (
	repositoryContract "github.com/dedekrnwan/go-clean/internal/contract/repository"
	usecaseContract "github.com/dedekrnwan/go-clean/internal/contract/usecase"
)

type Contract struct {
	Usecase    *usecaseContract.Contract
	Repository *repositoryContract.Contract
}

func New() *Contract {
	c := &Contract{}
	c.Repository = repositoryContract.New()
	c.Usecase = usecaseContract.New(c.Repository)
	return c
}
