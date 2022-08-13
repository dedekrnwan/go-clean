package response

import (
	"github.com/dedekrnwan/go-clean/model"
)

type Meta struct {
	Success bool                  `json:"success" default:"true"`
	Message string                `json:"message" default:"true"`
	Info    *model.PaginationInfo `json:"info"`
}

type responseHelper struct {
	Error   errorHelper
	Success successHelper
}

var Constant responseHelper = responseHelper{
	Error:   errorConstant,
	Success: successConstant,
}
