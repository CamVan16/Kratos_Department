package data

import (
	"DM/internal/biz"
	"DM/internal/models"
	"context"

	"gorm.io/gorm"
)

type departmentRepo struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *Data) biz.DepartmentRepo {
	return &departmentRepo{db: db.DB}
}

func (r *departmentRepo) Create(ctx context.Context, dept *models.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepo) GetByID(ctx context.Context, id uint32) (*models.Department, error) {
	var dept models.Department
	err := r.db.First(&dept, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepo) Update(ctx context.Context, dept *models.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&models.Department{}, id).Error
}

func (r *departmentRepo) GetAll(ctx context.Context) ([]*models.Department, error) {
	var depts []*models.Department
	err := r.db.Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}
