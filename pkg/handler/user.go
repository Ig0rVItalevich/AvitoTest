package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getUserBalance
// @Summary Get user balance
// @Description User balance receipt method
// @Produce json
// @Param id path int true "id of username"
// @Success 200 {object} models.User "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/users/{id} [GET]
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

// refillUserBalance
// @Summary Refill user balance
// @Description The method of accruing funds to the balance
// @Accept json
// @Produce json
// @Param input body models.RefillBalance true "refill balance struct"
// @Success 200 {object} handler.statusResponse "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 422 {object} handler.errorResponse "invalid request data"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/users/refill [POST]
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

// transfer
// @Summary Transfer
// @Description Method of transferring funds between users
// @Accept json
// @Produce json
// @Param input body models.Transfer true "transfer struct"
// @Success 200 {object} handler.statusResponse "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 422 {object} handler.errorResponse "insufficient funds to transfer"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/transfer [POST]
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
