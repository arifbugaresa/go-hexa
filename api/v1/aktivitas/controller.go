package aktivitas

import (
	"github.com/arifbugaresa/go-hexa/api/v1/aktivitas/dto"
	"github.com/arifbugaresa/go-hexa/api/v1/common"
	"github.com/arifbugaresa/go-hexa/business/aktivitas"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Controller struct {
	service aktivitas.Service
}

func NewController(service aktivitas.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateAktivitas(context echo.Context) (err error) {
	userIDLoggedIn := int64(context.Get("currentUser").(int))
	var request dto.AktivitasRequest
	if err = context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	request.UserInfoId = userIDLoggedIn

	err = c.service.CreateAktivitas(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Berhasil membuat aktivitas"))
}

func (c *Controller) UpdateAktivitas(context echo.Context) (err error) {
	userIDLoggedIn := int64(context.Get("currentUser").(int))
	var request dto.AktivitasRequest
	if err = context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	// mengambil id dari url param
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	request.ID = int64(id)
	request.UserInfoId = userIDLoggedIn

	err = c.service.UpdateAktivitas(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Berhasil mengubah aktivitas"))

}

func (c *Controller) DeleteAktivitas(context echo.Context) (err error) {
	userIDLoggedIn := int64(context.Get("currentUser").(int))
	var request dto.AktivitasRequest

	// mengambil id dari url param
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	request.ID = int64(id)
	request.UserInfoId = userIDLoggedIn

	err = c.service.DeleteAktivitas(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Berhasil menghapus aktivitas"))

}

func (c *Controller) GetListAktivitas(context echo.Context) (err error) {
	userIDLoggedIn := context.Get("currentUser").(int)

	listAktivitas, err := c.service.GetListAktivitas(userIDLoggedIn)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithData("Berhasil mengambil daftar aktivitas", listAktivitas))

}
