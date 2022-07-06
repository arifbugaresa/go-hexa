package user_info

import (
	"github.com/arifbugaresa/go-hexa/api/v1/common"
	"github.com/arifbugaresa/go-hexa/api/v1/user_info/dto"
	"github.com/arifbugaresa/go-hexa/business/user_info"
	"github.com/arifbugaresa/go-hexa/middleware/auth"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service     user_info.Service
	authService auth.Service
}

func NewController(service user_info.Service, authService auth.Service) *Controller {
	return &Controller{
		service:     service,
		authService: authService,
	}
}

func (c *Controller) Login(context echo.Context) (err error) {
	var request dto.UserLoginRequest
	if err = context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	userOnDB, err := c.service.Login(request)
	if err != nil {
		return context.JSON(common.NewBadRequestEmailOrPassword())
	}

	// generate token
	token, err := c.authService.GenerateToken(userOnDB.ID)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithData("Success Login.", dto.UserLoginResponse{Token: token}))
}

func (c *Controller) Logout(context echo.Context) (err error) {
	var request dto.UserLogoutRequest
	if err = context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	err = c.service.Logout(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData("Success Logout."))
}

func (c *Controller) GetListUserInfo(context echo.Context) (err error) {
	var response []dto.GetListUserResponse

	response, err = c.service.GetListUserInfo()
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithData("Success Get List Data", response))
}
