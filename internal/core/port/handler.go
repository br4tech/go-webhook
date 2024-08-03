package port

import "github.com/labstack/echo/v4"

type ISubscribeHandler interface {
	Create(c echo.Context) error
}

type IPaymentHandler interface {
	FindByOrdeId(c echo.Context) error
	Create(c echo.Context) error
}

type IProductHandler interface {
	Create(c echo.Context) error
}

type IOrderHandler interface {
	Create(c echo.Context) error
}
