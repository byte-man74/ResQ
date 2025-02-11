package repository

import (
	"example.com/resq/server/internal/models"
	"gorm.io/gorm"
)


type IncidentRepository interface {
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	ViewIncident(id string) (*models.Incident, error)
	EditIncident(id string, incident *models.Incident) (*models.Incident, error)
}

type incidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) IncidentRepository {
	return &incidentRepository{db: db}
}



func (r *incidentRepository) CreateIncident (incident *models.Incident) (*models.Incident, error) {
	result := r.db.Create(incident)

	if result.Error != nil {
		return nil, result.Error
	}

	return incident, nil
}

func (r *incidentRepository) ViewIncident (id string) (*models.Incident, error) {
	var incident models.Incident

	result := r.db.First(&incident, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &incident, nil
}

func (r *incidentRepository) EditIncident (id string, incident *models.Incident ) (*models.Incident, error ) {
	result, err := r.ViewIncident(id)

	if err != nil {
		return nil, err
	}

	updateResult := r.db.Model(&result).Updates(incident)
	if updateResult.Error != nil {
		return nil, updateResult.Error
	}

	return result, nil

}
