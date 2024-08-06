package biz

import (
	"DM/internal/entity"
	"context"
)

type SubDepartmentRepo interface {
	Create(ctx context.Context, dept *entity.SubDepartment) error
	GetByID(ctx context.Context, id uint32) (*entity.SubDepartment, error)
	Update(ctx context.Context, dept *entity.SubDepartment) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*entity.SubDepartment, error)
	GetByPage(ctx context.Context, page, limit uint32) ([]*entity.SubDepartment, error)
}

type SubDepartmentUC struct {
	repo SubDepartmentRepo
}

func NewSubDepartmentUC(repo SubDepartmentRepo) *SubDepartmentUC {
	return &SubDepartmentUC{repo: repo}
}

func (uc *SubDepartmentUC) CreateSubDepartment(ctx context.Context, name string, department_id uint32) (uint32, error) {
	subdept := &entity.SubDepartment{Name: name, DepartmentID: department_id}
	if err := uc.repo.Create(ctx, subdept); err != nil {
		return 0, err
	}
	return subdept.ID, nil
}

func (uc *SubDepartmentUC) GetSubDepartmentByID(ctx context.Context, id uint32) (*entity.SubDepartment, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *SubDepartmentUC) GetAllSubDepartment(ctx context.Context) ([]*entity.SubDepartment, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *SubDepartmentUC) UpdateSubDepartment(ctx context.Context, id uint32, name string, department_id uint32) error {
	dept := &entity.SubDepartment{ID: id, Name: name, DepartmentID: department_id}
	return uc.repo.Update(ctx, dept)
}

func (uc *SubDepartmentUC) DeleteSubDepartment(ctx context.Context, id uint32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *SubDepartmentUC) GetSubDepartmentByPage(ctx context.Context, page, limit uint32) ([]*entity.SubDepartment, error) {
	return uc.repo.GetByPage(ctx, page, limit)
}
