package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/yogapratama23/test-gotempl/helpers"
	"github.com/yogapratama23/test-gotempl/models"
	"github.com/yogapratama23/test-gotempl/views/user"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	userData := models.User{
		Username: "john",
		Email:    "john@gmail.com",
	}
	return helpers.Render(c, user.Show(userData))
}
