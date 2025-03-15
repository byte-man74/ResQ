package report

type ReportService interface {
}

type reportService struct {
	repository ReportRepository
}

func NewReportService(repo *ReportRepository) ReportService {
	return &reportService{repository: repo}
}


	