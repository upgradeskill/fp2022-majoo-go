package user

import (
	"mini-pos/core/user/entity"
	"mini-pos/handlers/util"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInsertRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateByIdRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserDefaultResponse struct {
	Id       uint      `json:"id"`
	Email    string    `json:"email"`
	IsAdmin  bool      `json:"is_admin"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func UserNewDefaultResponse(data *entity.User) *UserDefaultResponse {
	return &UserDefaultResponse{
		Id:       data.Id,
		Email:    data.Email,
		IsAdmin:  data.IsAdmin,
		Created:  data.Created,
		Modified: data.Modified,
	}
}

func (handler Handler) UserList(c echo.Context) error {
	queryParams := c.QueryParams()
	cleanQueryParams := util.QueryParamsCleaner(queryParams)
	result, err := handler.service.UserList(cleanQueryParams.QueryParams, cleanQueryParams.PerPage, cleanQueryParams.Offset)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	var results []UserDefaultResponse
	for _, data := range result {
		results = append(results, *UserNewDefaultResponse(data))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], results))
}

func (handler Handler) UserCreate(c echo.Context) error {
	bodyRequest := new(UserInsertRequest)
	if err := c.Bind(bodyRequest); err != nil {
		result := util.Map["badRequest"]
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	if err := c.Validate(bodyRequest); err != nil {
		errors := util.BuildErrorBodyRequestValidator(err)
		result := util.Map["badRequest"]
		result.Errors = errors
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	data := entity.User{
		Email:    bodyRequest.Email,
		Password: bodyRequest.Password,
	}
	result, err := handler.service.UserCreate(data)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusCreated, util.NewResponse("", util.Map["created"], UserNewDefaultResponse(result)))
}

func (handler Handler) UserFindById(c echo.Context) error {
	id := c.Param("id")
	result, err := handler.service.UserFindById(id)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], UserNewDefaultResponse(result)))
}

func (handler Handler) UserUpdateById(c echo.Context) error {
	id := c.Param("id")
	bodyRequest := new(UserUpdateByIdRequest)
	if err := c.Bind(bodyRequest); err != nil {
		result := util.Map["badRequest"]
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	if err := c.Validate(bodyRequest); err != nil {
		errors := util.BuildErrorBodyRequestValidator(err)
		result := util.Map["badRequest"]
		result.Errors = errors
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	data := entity.User{
		Email:    bodyRequest.Email,
		Password: bodyRequest.Password,
		IsAdmin:  bodyRequest.IsAdmin,
	}
	result, err := handler.service.UserUpdateById(id, data)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = []string{err.Error()}
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], UserNewDefaultResponse(result)))
}

func (handler Handler) UserDeleteById(c echo.Context) error {
	id := c.Param("id")
	err := handler.service.UserDeleteById(id)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = []string{err.Error()}
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["deleted"], nil))
}
