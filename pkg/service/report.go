package service

import (
	"encoding/csv"
	"fmt"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/repository"
	"os"
)

type ReportService struct {
	repos repository.Report
}

var _ Report = (*ReportService)(nil)

func NewReportService(repos repository.Report) *ReportService {
	return &ReportService{repos: repos}
}

func (s *ReportService) GetUserReport(input models.InputUserReport) ([]models.HistoryRow, error) {
	switch input.OrderBy {
	case "sum":
		input.OrderBy = "amount"
	case "date":
		input.OrderBy = "accepted_date"
	default:
		input.OrderBy = "accepted_date"
	}

	return s.repos.GetUserReport(input)
}

func (s *ReportService) GetRevenueReport(year int, month int) (string, error) {
	revenueReport, err := s.repos.GetRevenueReport(year, month)
	if err != nil {
		return "", err
	}

	if err := checkReportsDir(); err != nil {
		return "", err
	}

	path := fmt.Sprintf("./reports/report-%d-%d.csv", year, month)
	csvFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	for _, record := range revenueReport {
		err = writer.Write([]string{fmt.Sprintf("%d", record.ProductId), fmt.Sprintf("%f", record.Sum)})
		if err != nil {
			break
		}
	}
	writer.Flush()

	return path, nil
}

func checkReportsDir() error {
	info, err := os.Stat("./reports")
	if os.IsNotExist(err) || !info.IsDir() {
		if err := os.Mkdir("reports", 0777); err != nil {
			return err
		}
	}

	return nil
}
