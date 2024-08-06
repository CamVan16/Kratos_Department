package biz

import (
	"DM/internal/entity"
	"context"
)

type DepartmentRepo interface {
	Create(ctx context.Context, dept *entity.Department) error
	GetByID(ctx context.Context, id uint32) (*entity.Department, error)
	Update(ctx context.Context, dept *entity.Department) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*entity.Department, error)
	GetByPage(ctx context.Context, page, limit uint32) ([]*entity.Department, error)
}

type DepartmentUC struct {
	repo DepartmentRepo
}

func NewDepartmentUC(repo DepartmentRepo) *DepartmentUC {
	return &DepartmentUC{repo: repo}
}

func (uc *DepartmentUC) CreateDepartment(ctx context.Context, name string) (uint32, error) {
	dept := &entity.Department{Name: name}
	if err := uc.repo.Create(ctx, dept); err != nil {
		return 0, err
	}
	return dept.ID, nil
}

func (uc *DepartmentUC) GetDepartmentByID(ctx context.Context, id uint32) (*entity.Department, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *DepartmentUC) UpdateDepartment(ctx context.Context, id uint32, name string) error {
	dept := &entity.Department{ID: id, Name: name}
	return uc.repo.Update(ctx, dept)
}

func (uc *DepartmentUC) DeleteDepartment(ctx context.Context, id uint32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *DepartmentUC) GetAllDepartment(ctx context.Context) ([]*entity.Department, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *DepartmentUC) GetDepartmentByPage(ctx context.Context, page, limit uint32) ([]*entity.Department, error) {
	return uc.repo.GetByPage(ctx, page, limit)
}
