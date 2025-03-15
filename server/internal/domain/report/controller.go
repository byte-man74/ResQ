package report

import (
	"github.com/gin-gonic/gin"
)

type ReportController interface {
}

type reportController struct {
	service ReportService
}

func NewReportController(service ReportService) ReportController {
	return &reportController{service: service}
}
