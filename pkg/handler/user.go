package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserBalance(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.Get(int(id))
	if err != nil {
		switch err {
		case service.ErrUserDoesNotExist:
			NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) refillUserBalance(ctx *gin.Context) {
	var input models.RefillBalance
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusUnprocessableEntity, service.ErrIncorrectParameters.Error())
		return
	}

	err := h.services.User.Refill(input.UserId, input.Amount)
	if err != nil {
		switch err {
		case service.ErrUserDoesNotExist:
			NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
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

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, service.ErrIncorrectParameters.Error())
		return
	}

	err := h.services.User.ExecuteTransfer(input)
	if err != nil {
		switch err {
		case service.ErrNotEnoughMoney:
			NewErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		case service.ErrUserDoesNotExist:
			NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
