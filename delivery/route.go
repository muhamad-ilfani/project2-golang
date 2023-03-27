package delivery

import (
	"context"

	"github.com/labstack/echo"
)

func (x *echoObject) initRoute(ctx context.Context) {
	x.Echo = echo.New()
	x.Echo.GET("/", welcome(ctx))
	x.Echo.POST("/login", RegisterUser(ctx, x.UserUseCase))
}