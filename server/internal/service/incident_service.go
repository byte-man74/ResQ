package service

import (
	"fmt"

	"resq.com/resq/server/models"
	"resq.com/resq/server/internal/repository"
)

type IncidentService interface {
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	UpdateIncident(id string, incident *models.Incident) (*models.Incident, error)
	ViewIncident(id string) (*models.Incident, error)
}

type incidentService struct {
	repo repository.IncidentRepository
}

func NewIncidentService(repo repository.IncidentRepository) IncidentService {
	return &incidentService{repo: repo}
}

func (i *incidentService) CreateIncident(incident *models.Incident) (*models.Incident, error) {
	createdIncident, err := i.repo.CreateIncident(incident)

	if err != nil {
		return nil, fmt.Errorf("unable to create incident %w", err)
	}

	return createdIncident, nil
}

func (i *incidentService) UpdateIncident(id string, incident *models.Incident) (*models.Incident, error) {
	updatedIncident, err := i.repo.UpdateIncident(id, incident)

	if err != nil {
		return nil, fmt.Errorf("unable to update incident %w", err)
	}

	return updatedIncident, nil
}

func (i *incidentService) ViewIncident(id string) (*models.Incident, error) {
	incident, err := i.repo.ViewIncident(id)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch incident %w", err)
	}

	return incident, nil
}
