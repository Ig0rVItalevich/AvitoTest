package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reservePurchase
// @Summary Reserve purchase
// @Description Method of reserving funds from the main balance in a separate account
// @Accept json
// @Produce json
// @Param input body models.Purchase true "purchase struct"
// @Success 200 {object} handler.statusResponse "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 422 {object} handler.errorResponse "not enough money to buy"
// @Failure 424 {object} handler.errorResponse "this service is already reserved"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/purchase/reserve [POST]
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

// acceptPurchase
// @Summary Accept purchase
// @Description Revenue recognition method
// @Accept json
// @Produce json
// @Param input body models.Purchase true "purchase struct"
// @Success 200 {object} handler.statusResponse "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 424 {object} handler.errorResponse "reserved service not found"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/purchase/accept [POST]
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

// cancelPurchase
// @Summary Cancel purchase
// @Description Reservation Method
// @Accept json
// @Produce json
// @Param input body models.Purchase true "purchase struct"
// @Success 200 {object} handler.statusResponse "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 424 {object} handler.errorResponse "reserved service not found"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/purchase/cancel [POST]
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
