package handler

import (
	"net/http"
	"unicode/utf8"

	"github.com/ds124wfegd/mp-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input mp.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if utf8.RuneCountInString(input.Username) < 8 || utf8.RuneCountInString(input.Username) > 20 {
		newErrorResponse(c, http.StatusBadRequest, "the username must contain from 8 to 20 characterss")
	}

	if utf8.RuneCountInString(input.Password) < 8 || utf8.RuneCountInString(input.Password) > 20 {
		newErrorResponse(c, http.StatusBadRequest, "the password must contain from 8 to 20 characterss")
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
