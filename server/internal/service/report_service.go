package service

import (
	"fmt"

	"resq.com/resq/server/models"
	"resq.com/resq/server/internal/repository"
)

type ReportService interface {
	CreateReport(report *models.Report) (*models.Report, error)
	EditReport(id string, report *models.Report) (*models.Report, error)
	ViewReport(id string) (*models.Report, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (r *reportService) CreateReport(report *models.Report) (*models.Report, error) {
	createdReport, err := r.repo.CreateReport(report)

	if err != nil {
		return nil, fmt.Errorf("unable to create report %w", err)
	}

	return createdReport, nil
}

func (r *reportService) EditReport(id string, report *models.Report) (*models.Report, error) {
	editedReport, err := r.repo.EditReport(id, report)

	if err != nil {
		return nil, fmt.Errorf("unable to edit report %w", err)
	}

	return editedReport, nil
}

func (r *reportService) ViewReport(id string) (*models.Report, error) {
	report, err := r.repo.ViewReport(id)

	if err != nil {
		return nil, fmt.Errorf("unable to view report %w", err)
	}

	return report, nil
}
