package port

import "github.com/labstack/echo/v4"

type ISubscriptionHandler interface {
	Create(c echo.Context) error
}
