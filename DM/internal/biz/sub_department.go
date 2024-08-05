package biz

import (
	"DM/internal/models"
	"context"
)

type SubDepartmentRepo interface {
	Create(ctx context.Context, dept *models.SubDepartment) error
	GetByID(ctx context.Context, id uint32) (*models.SubDepartment, error)
	Update(ctx context.Context, dept *models.SubDepartment) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*models.SubDepartment, error)
}

type SubDepartmentService struct {
	repo SubDepartmentRepo
}

func NewSubDepartmentService(repo SubDepartmentRepo) *SubDepartmentService {
	return &SubDepartmentService{repo: repo}
}

func (s *SubDepartmentService) CreateSubDepartment(ctx context.Context, name string, department_id uint32) (uint32, error) {
	subdept := &models.SubDepartment{Name: name, DepartmentID: department_id}
	if err := s.repo.Create(ctx, subdept); err != nil {
		return 0, err
	}
	return subdept.ID, nil
}

func (s *SubDepartmentService) GetSubDepartmentByID(ctx context.Context, id uint32) (*models.SubDepartment, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SubDepartmentService) GetAllSubDepartment(ctx context.Context) ([]*models.SubDepartment, error) {
	return s.repo.GetAll(ctx)
}

// func (s *subDepartmentService) GetSubDepartmentsByDepartmentID(departmentID uint) ([]models.SubDepartment, error) {
// 	return s.repository.FindByDepartmentID(departmentID)
// }

func (s *SubDepartmentService) UpdateSubDepartment(ctx context.Context, id uint32, name string, department_id uint32) error {
	dept := &models.SubDepartment{ID: id, Name: name, DepartmentID: department_id}
	return s.repo.Update(ctx, dept)
}

func (s *SubDepartmentService) DeleteSubDepartment(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}
