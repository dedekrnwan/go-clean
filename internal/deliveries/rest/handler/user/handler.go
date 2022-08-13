package user

import (
	"encoding/json"
	"fmt"
	"github.com/dedekrnwan/go-clean/internal/factory"
	"github.com/dedekrnwan/go-clean/internal/usecase"
	"github.com/dedekrnwan/go-clean/model"
	"github.com/dedekrnwan/go-clean/pkg/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseUser usecase.User
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		usecaseUser: f.Usecase.User,
	}
}

func (h *handler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	payload := model.NewHttpQuery(c.Request(), model.User{})
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "testing user failed")
	}
	payload.BindFilters()
	// if err := c.Validate(payload); err != nil {
	// 	fmt.Println(err.Error())
	// 	return c.String(http.StatusBadRequest, "testing user failed")
	// }
	// for _, v := range payload.Filters {
	// 	fmt.Printf("%s %s %s\n", v.Field, v.Operator, v.Value)
	// }
	// for _, v := range payload.Ascending {
	// 	fmt.Printf("%s n", v)
	// }
	fmt.Println(payload.Search)
	// fmt.Println(payload.Ascending)

	data, info, err := h.usecaseUser.Find(ctx, payload.Search, payload.Filters, payload.Ascending, payload.Descending, payload.Pagination, payload.Preloads)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	//transform
	bytes, err := json.Marshal(data)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	result := []model.User{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.CustomSuccessBuilder(response.Constant.Success.OK.Code, result, "Data has been retrieve", info).Send(c)
}
