package absensi

import (
	"github.com/arifbugaresa/go-hexa/api/v1/common"
	"github.com/arifbugaresa/go-hexa/business/absensi"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service absensi.Service
}

func NewController(service absensi.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CheckIn(context echo.Context) (err error) {
	userIDLoggedIn := context.Get("currentUser").(int)

	err = c.service.CheckIn(userIDLoggedIn)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Berhasil CheckIn"))
}

func (c *Controller) CheckOut(context echo.Context) (err error) {
	userIDLoggedIn := context.Get("currentUser").(int)

	err = c.service.CheckOut(userIDLoggedIn)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Berhasil CheckOut"))
}

func (c *Controller) GetListAbsensiByIDUser(context echo.Context) (err error) {
	userIDLoggedIn := context.Get("currentUser").(int)

	listAbsensi, err := c.service.GetListAbsensi(userIDLoggedIn)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithData("Berhasil Mengambil Data Absensi", listAbsensi))
}
