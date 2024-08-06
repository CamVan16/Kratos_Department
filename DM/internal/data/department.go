package data

import (
	"DM/internal/biz"
	"DM/internal/entity"
	"context"

	"gorm.io/gorm"
)

type departmentRepo struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *Data) biz.DepartmentRepo {
	return &departmentRepo{db: db.DB}
}

func (r *departmentRepo) Create(ctx context.Context, dept *entity.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepo) GetByID(ctx context.Context, id uint32) (*entity.Department, error) {
	var dept entity.Department
	err := r.db.First(&dept, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepo) Update(ctx context.Context, dept *entity.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&entity.Department{}, id).Error
}

func (r *departmentRepo) GetAll(ctx context.Context) ([]*entity.Department, error) {
	var depts []*entity.Department
	err := r.db.Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}

func (r *departmentRepo) GetByPage(ctx context.Context, page, limit uint32) ([]*entity.Department, error) {
	var depts []*entity.Department
	if int(page) <= 0 {
		page = 1
	}
	err := r.db.Offset(int(limit * (page - 1))).Limit(int(limit)).Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}
