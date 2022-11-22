package controller

import (
	"github.com/DendiAnugerah/Todo/model"
	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	var user model.User

	err := c.
}
