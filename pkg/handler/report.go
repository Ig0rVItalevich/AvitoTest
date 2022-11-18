package handler

import (
	"fmt"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHistory struct {
	History []models.HistoryRow `json:"history"`
}

func (h *Handler) getReportUser(ctx *gin.Context) {
	var input models.InputUserReport
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !input.Validate() {
		fmt.Println(input)
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect parameters")
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

func (h *Handler) getReportRevenue(ctx *gin.Context) {
	var input inputRevenueReport
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.Month < january || input.Month > december || input.Year < yearBottom || input.Year > yearHigh {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect parameters")
		return
	}

	path, err := h.services.Report.GetRevenueReport(input.Year, input.Month)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"file_path": path,
	})
}
