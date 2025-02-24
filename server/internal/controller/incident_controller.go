package controller

import (
	"example.com/resq/server/internal/service"
	"github.com/gin-gonic/gin"
)


type IncidentController interface {
	CreateIncident (ctx *gin.Context)
}



type incidentController struct  {
	service service.IncidentService
}


func NewIncidentController (service service.IncidentService) IncidentController {
	return &incidentController{service: service}
}


func (r *incidentController) CreateIncident (ctx *gin.Context) {
	
}
