package handler

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHistory struct {
	History []models.HistoryRow `json:"history"`
}

// getReportUser
// @Summary Get user report
// @Description Method for obtaining a list of transactions with comments from where and why the funds were credited / debited from the balance
// @Accept json
// @Produce json
// @Param input body models.InputUserReport true "input user report struct"
// @Success 200 {object} userHistory "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 401 {object} handler.errorResponse "the user with the given id does not exist"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/reports/user [GET]
func (h *Handler) getReportUser(ctx *gin.Context) {
	var input models.InputUserReport
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect parameters")
		return
	}

	flag, err := h.services.User.Exist(input.UserId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if !flag {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	history, err := h.services.Report.GetUserReport(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, userHistory{
		History: history,
	})
}

type inputRevenueReport struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

func (i *inputRevenueReport) Validate() bool {
	if i.Month < january || i.Month > december || i.Year < yearBottom || i.Year > yearHigh {
		return false
	}

	return true
}

type pathOutput struct {
	FilePath string `json:"file_path"`
}

// getReportRevenue
// @Summary Get revenue report
// @Description Monthly report receipt method
// @Accept json
// @Produce json
// @Param input body inputRevenueReport true "input revenue report struct"
// @Success 200 {object} pathOutput "OK"
// @Failure 400 {object} handler.errorResponse "invalid request parameters"
// @Failure 500 {object} handler.errorResponse "server error"
// @Router /api/v1/reports/revenue [GET]
func (h *Handler) getReportRevenue(ctx *gin.Context) {
	var input inputRevenueReport
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect parameters")
		return
	}

	path, err := h.services.Report.GetRevenueReport(input.Year, input.Month)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, pathOutput{
		FilePath: path,
	})
}
