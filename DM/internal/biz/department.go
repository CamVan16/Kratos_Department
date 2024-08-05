package biz

import (
	"DM/internal/models"
	"context"
)

type DepartmentRepo interface {
	Create(ctx context.Context, dept *models.Department) error
	GetByID(ctx context.Context, id uint32) (*models.Department, error)
	Update(ctx context.Context, dept *models.Department) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*models.Department, error)
}

type DepartmentService struct {
	repo DepartmentRepo
}

func NewDepartmentService(repo DepartmentRepo) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, name string) (uint32, error) {
	dept := &models.Department{Name: name}
	if err := s.repo.Create(ctx, dept); err != nil {
		return 0, err
	}
	return dept.ID, nil
}

func (s *DepartmentService) GetDepartmentByID(ctx context.Context, id uint32) (*models.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, id uint32, name string) error {
	dept := &models.Department{ID: id, Name: name}
	return s.repo.Update(ctx, dept)
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}

func (s *DepartmentService) GetAllDepartment(ctx context.Context) ([]*models.Department, error) {
	return s.repo.GetAll(ctx)
}
