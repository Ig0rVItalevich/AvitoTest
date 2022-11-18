package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserBalance(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.Get(int(id))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type refillBalance struct {
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *Handler) refillUserBalance(ctx *gin.Context) {
	var input refillBalance
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.User.Refill(input.UserId, input.Amount)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) transfer(ctx *gin.Context) {
	var input models.Transfer
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.Amount <= 0 {
		NewErrorResponse(ctx, http.StatusBadRequest, "wrong amount of money")
		return
	}

	err := h.services.User.ExecuteTransfer(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
