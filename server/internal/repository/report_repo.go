package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"resq.com/resq/server/internal/models"
)

type ReportRepository interface {
	CreateReport(report *models.Report) (*models.Report, error)
	EditReport(id string, report *models.Report) (*models.Report, error)
	ViewReport(id string) (*models.Report, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) CreateReport(report *models.Report) (*models.Report, error) {
	if report == nil {
		return nil, errors.New("report cannot be nil")
	}

	result := r.db.Create(report)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to create report: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no report was created")
	}

	return report, nil
}

func (r *reportRepository) ViewReport(id string) (*models.Report, error) {
	var report models.Report
	result := r.db.First(&report, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &report, nil
}

func (r *reportRepository) EditReport(id string, report *models.Report) (*models.Report, error) {
	if report == nil {
		return nil, errors.New("report cannot be nil")
	}

	existingReport, err := r.ViewReport(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find report: %w", err)
	}

	result := r.db.Model(&existingReport).Updates(report)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update report: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no report was updated")
	}

	return report, nil
}
