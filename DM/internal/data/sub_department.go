package data

import (
	"DM/internal/biz"
	"DM/internal/models"
	"context"

	"gorm.io/gorm"
)

type subDepartmentRepo struct {
	db *gorm.DB
}

func NewSubDepartmentRepository(db *Data) biz.SubDepartmentRepo {
	return &subDepartmentRepo{db: db.DB}
}

func (r *subDepartmentRepo) Create(ctx context.Context, subdept *models.SubDepartment) error {
	return r.db.Create(subdept).Error
}

func (r *subDepartmentRepo) GetByID(ctx context.Context, id uint32) (*models.SubDepartment, error) {
	var subdept models.SubDepartment
	err := r.db.First(&subdept, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &subdept, nil
}

func (r *subDepartmentRepo) Update(ctx context.Context, subdept *models.SubDepartment) error {
	return r.db.Save(subdept).Error
}

func (r *subDepartmentRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&models.SubDepartment{}, id).Error
}

func (r *subDepartmentRepo) GetAll(ctx context.Context) ([]*models.SubDepartment, error) {
	var subdepts []*models.SubDepartment
	err := r.db.Find(&subdepts).Error
	if err != nil {
		return nil, err
	}
	return subdepts, nil
}
