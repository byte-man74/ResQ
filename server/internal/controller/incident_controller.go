package controller

import (
	"github.com/gin-gonic/gin"
	"resq.com/resq/server/internal/service"
)

type IncidentController interface {
	CreateIncident(ctx *gin.Context)
}

type incidentController struct {
	service service.IncidentService
}

func NewIncidentController(service service.IncidentService) IncidentController {
	return &incidentController{service: service}
}

func (r *incidentController) CreateIncident(ctx *gin.Context) {

}
