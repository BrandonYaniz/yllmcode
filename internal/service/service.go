package service

import (
	"context"

	"github.com/BrandonYaniz/yllmcode/internal/project"
)

type Service struct{}

type InitProjectRequest = project.InitProjectRequest
type InitProjectResult = project.InitProjectResult

func New() *Service {
	return &Service{}
}

func (s *Service) InitProject(ctx context.Context, req InitProjectRequest) (*InitProjectResult, error) {
	return project.InitProject(ctx, req)
}
