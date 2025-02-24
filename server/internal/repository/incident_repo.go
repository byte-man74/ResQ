package repository

import (
	"fmt"
	"example.com/resq/server/internal/models"
	"gorm.io/gorm"
)



type IncidentRepository interface {
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	UpdateIncident(id string, incident *models.Incident) (*models.Incident, error)
}


type incidentRepository struct {
	db *gorm.DB
}


func NewIncidentRepository (db *gorm.DB) IncidentRepository {
	return &incidentRepository{db: db}
}


func (r *incidentRepository) CreateIncident(incident *models.Incident) (*models.Incident, error) {
	createdIncident := r.db.Create(incident)

	if createdIncident.Error != nil {
		return nil, fmt.Errorf("failed to create incident %w", createdIncident.Error)
	}

	return incident, nil
}


func (r *incidentRepository) UpdateIncident(id string, incident *models.Incident) (*models.Incident, error) {

	if id == "" {
		return nil, fmt.Errorf("id cannot be nil")
	}

	result := r.db.First(incident, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get the incident by id: %w", result.Error)
	}

	result = r.db.Save(incident)
	if result.Error != nil {
		return nil, fmt.Errorf("error updating incident: %w", result.Error)
	}

	return incident, nil
}

func (r *incidentRepository) DeleteIncident(id string) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	result := r.db.Delete(&models.Incident{}, "id = ?", id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete incident: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no incident was deleted")
	}

	return nil
}



