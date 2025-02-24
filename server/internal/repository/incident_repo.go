package repository

import (
	"fmt"
	"example.com/resq/server/internal/models"
	"gorm.io/gorm"
)

// Interface Definition
type IncidentRepository interface {
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	UpdateIncident(id string, incident *models.Incident) (*models.Incident, error)
	ViewIncident(id string) (*models.Incident, error)
	DeleteIncident(id string) error
}

// Repository Implementation
type incidentRepository struct {
	db *gorm.DB
}

// Constructor for Dependency Injection
func NewIncidentRepository(db *gorm.DB) IncidentRepository {
	return &incidentRepository{db: db}
}

// CreateIncident inserts a new incident into the database
func (r *incidentRepository) CreateIncident(incident *models.Incident) (*models.Incident, error) {
	if err := r.db.Create(incident).Error; err != nil {
		return nil, fmt.Errorf("failed to create incident: %w", err)
	}
	return incident, nil
}

// UpdateIncident updates an existing incident based on ID
func (r *incidentRepository) UpdateIncident(id string, updatedIncident *models.Incident) (*models.Incident, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	var existingIncident models.Incident
	if err := r.db.First(&existingIncident, id).Error; err != nil {
		return nil, fmt.Errorf("incident not found: %w", err)
	}

	if err := r.db.Model(&existingIncident).Updates(updatedIncident).Error; err != nil {
		return nil, fmt.Errorf("failed to update incident: %w", err)
	}

	return &existingIncident, nil
}

// DeleteIncident deletes an incident by ID
func (r *incidentRepository) DeleteIncident(id string) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	result := r.db.Delete(&models.Incident{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete incident: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no incident found to delete")
	}

	return nil
}

// ViewIncident retrieves an incident by ID
func (r *incidentRepository) ViewIncident(id string) (*models.Incident, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	var incident models.Incident
	if err := r.db.First(&incident, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("incident not found")
		}
		return nil, fmt.Errorf("failed to fetch incident: %w", err)
	}

	return &incident, nil
}
