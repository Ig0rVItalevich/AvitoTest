package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reservePurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.Get(input.UserId)
	if user.Balance-input.Amount < 0 {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Purchase.Reserve(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) acceptPurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.Purchase.Accept(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) cancelPurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.Purchase.Cancel(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
