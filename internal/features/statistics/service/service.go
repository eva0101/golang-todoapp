package statistics_service

import (
	"context"
	"time"

	"github.com/eva0101/golang-todoapp/internal/core/domain"
)

type StaitsticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(
	statisticsRepository StatisticsRepository,
) *StaitsticsService {
	return &StaitsticsService{
		statisticsRepository: statisticsRepository,
	}
}
