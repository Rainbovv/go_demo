package handler

import (
	"github.com/gin-gonic/gin"
	"go-demo/types"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input types.User

	if err := ctx.BindJSON(&input); err != nil {
		types.NeErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.AuthorizationService.CreateUser(input)
	if err != nil {
		types.NeErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		types.NeErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.AuthorizationService.SignIn(input.Username, input.Password)
	if err != nil {
		types.NeErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
