package outlet

import (
	"mini-pos/core/outlet/entity"
	"mini-pos/handlers/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OutletInsertRequest struct {
	Name string `json:"name" validate:"required"`
}

type OutletUserInsertRequest struct {
	OutletId int `json:"outlet_id" validate:"required"`
	UserId   int `json:"user_id" validate:"required"`
}

type OutletUpdatRequest struct {
	Name string `json:"name" validate:"required"`
}

type OutletUpdateByIdRequest struct {
	Id uint `json:"id" validate:"required"`
}

type OutletUserUpdateByIdRequest struct {
	OutletId int `json:"outlet_id" validate:"required"`
	UserId   int `json:"user_id" validate:"required"`
}

type OutletDefaultResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"email"`
}

func OutletNewDefaultResponse(data *entity.Outlet) *OutletDefaultResponse {
	return &OutletDefaultResponse{
		Id:   data.Id,
		Name: data.Name,
	}
}

func (handler Handler) OutletList(c echo.Context) error {
	queryParams := c.QueryParams()
	cleanQueryParams := util.QueryParamsCleaner(queryParams)
	result, err := handler.service.OutletList(cleanQueryParams.QueryParams, cleanQueryParams.PerPage, cleanQueryParams.Offset)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	var results []OutletDefaultResponse
	for _, data := range result {
		results = append(results, *OutletNewDefaultResponse(data))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], results))
}

func (handler Handler) OutletCreate(c echo.Context) error {
	bodyRequest := new(OutletInsertRequest)
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
	data := entity.Outlet{
		Name: bodyRequest.Name,
	}
	result, err := handler.service.OutletCreate(data)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusCreated, util.NewResponse("", util.Map["created"], OutletNewDefaultResponse(result)))
}

func (handler Handler) OutletFindById(c echo.Context) error {
	id := c.Param("id")
	result, err := handler.service.OutletFindById(id)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = append(result.Errors, err.Error())
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], OutletNewDefaultResponse(result)))
}

func (handler Handler) OutletUpdateById(c echo.Context) error {
	id := c.Param("id")
	bodyRequest := new(OutletUpdateByIdRequest)
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
	data := entity.Outlet{
		Id: bodyRequest.Id,
	}
	result, err := handler.service.OutletUpdateById(id, data)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = []string{err.Error()}
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["ok"], OutletNewDefaultResponse(result)))
}

func (handler Handler) OutletDeleteById(c echo.Context) error {
	id := c.Param("id")
	err := handler.service.OutletDeleteById(id)
	if err != nil {
		result := util.Map["badRequest"]
		result.Errors = []string{err.Error()}
		return c.JSON(http.StatusBadRequest, util.NewResponse("", result, nil))
	}
	return c.JSON(http.StatusOK, util.NewResponse("", util.Map["deleted"], nil))
}
