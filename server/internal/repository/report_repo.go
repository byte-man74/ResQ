package repository

import (
	"fmt"

	"example.com/resq/server/internal/models"
	"gorm.io/gorm"
)


type ReportRepository interface {
	CreateReport(report *models.Report) (*models.Report, error)
	EditReport(id string, report *models.Report) (*models.Report, error)
	ViewReport(id string) (*models.Report, error)
	DeleteReport(id string) (error)
}

type reportRepository struct {
	db *gorm.DB
}


func NewReportRepository (db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}



func (r *reportRepository) CreateReport(report *models.Report) (*models.Report, error) {
	result := r.db.Create(report)
	if result.Error != nil {
		return nil, result.Error
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
	result, err := r.ViewReport(id)

	if err != nil {
		return nil, err
	}

	updateResult := r.db.Model(&result).Updates(report)
	if updateResult.Error != nil {
		return nil, updateResult.Error
	}

	return result, nil
}

func (r *reportRepository) DeleteReport(id string) error {
	if id == "" {
		return fmt.Errorf("invalid id provided")
	}

	report, err := r.ViewReport(id)
	if err != nil {
		return fmt.Errorf("failed to find report: %w", err)
	}

	if report == nil {
		return fmt.Errorf("report not found")
	}

	if err := r.db.Delete(report).Error; err != nil {
		return fmt.Errorf("failed to delete report: %w", err)
	}

	return nil
}
