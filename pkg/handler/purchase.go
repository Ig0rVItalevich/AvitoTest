package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reservePurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, service.ErrIncorrectParameters.Error())
		return
	}

	flag, err := h.services.User.Exist(input.UserId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if !flag {
		NewErrorResponse(ctx, http.StatusUnauthorized, service.ErrUserDoesNotExist.Error())
		return
	}

	user, err := h.services.User.Get(input.UserId)
	if user.Balance-input.Amount < 0 {
		NewErrorResponse(ctx, http.StatusUnprocessableEntity, service.ErrNotEnoughMoney.Error())
		return
	}

	err = h.services.Purchase.Reserve(input)
	if err != nil {
		switch err {
		case service.ErrReservedPurchaseAlreadyExists:
			NewErrorResponse(ctx, http.StatusFailedDependency, err.Error())
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

func (h *Handler) acceptPurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, service.ErrIncorrectParameters.Error())
		return
	}

	flag, err := h.services.User.Exist(input.UserId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if !flag {
		NewErrorResponse(ctx, http.StatusUnauthorized, service.ErrUserDoesNotExist.Error())
		return
	}

	err = h.services.Purchase.Accept(input)
	if err != nil {
		switch err {
		case service.ErrReservedPurchaseDoesNotExist:
			NewErrorResponse(ctx, http.StatusFailedDependency, err.Error())
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

func (h *Handler) cancelPurchase(ctx *gin.Context) {
	var input models.Purchase
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, service.ErrIncorrectParameters.Error())
		return
	}

	flag, err := h.services.User.Exist(input.UserId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if !flag {
		NewErrorResponse(ctx, http.StatusUnauthorized, service.ErrUserDoesNotExist.Error())
		return
	}

	err = h.services.Purchase.Cancel(input)
	if err != nil {
		switch err {
		case service.ErrReservedPurchaseDoesNotExist:
			NewErrorResponse(ctx, http.StatusFailedDependency, err.Error())
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
