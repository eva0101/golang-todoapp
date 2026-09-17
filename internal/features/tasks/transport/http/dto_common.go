package tasks_transport_http

import (
	"time"

	"github.com/eva0101/golang-todoapp/internal/core/domain"
)

type DTOTaskResponse struct {
	ID           int        `json:"id"             example:"15"`
	Version      int        `json:"version"        example:"3"`
	Title        string     `json:"title"          example:"Домашка"`
	Description  *string    `json:"description"    example:"Сделать до четверга дз по математике"`
	Completed    bool       `json:"completed"      example:"false"`
	CreatedAt    time.Time  `json:"created_at"     example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at"   example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"5"`
}

func taskDTOFromDomain(task domain.Task) DTOTaskResponse {
	return DTOTaskResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func taskDTOsFromDomains(tasks []domain.Task) []DTOTaskResponse {
	dtos := make([]DTOTaskResponse, len(tasks))

	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}

	return dtos
}
