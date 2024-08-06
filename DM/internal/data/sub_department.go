package data

import (
	"DM/internal/biz"
	"DM/internal/entity"
	"context"

	"gorm.io/gorm"
)

type subDepartmentRepo struct {
	db *gorm.DB
}

func NewSubDepartmentRepository(db *Data) biz.SubDepartmentRepo {
	return &subDepartmentRepo{db: db.DB}
}

func (r *subDepartmentRepo) Create(ctx context.Context, subdept *entity.SubDepartment) error {
	return r.db.Create(subdept).Error
}

func (r *subDepartmentRepo) GetByID(ctx context.Context, id uint32) (*entity.SubDepartment, error) {
	var subdept entity.SubDepartment
	err := r.db.First(&subdept, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &subdept, nil
}

func (r *subDepartmentRepo) Update(ctx context.Context, subdept *entity.SubDepartment) error {
	return r.db.Save(subdept).Error
}

func (r *subDepartmentRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&entity.SubDepartment{}, id).Error
}

func (r *subDepartmentRepo) GetAll(ctx context.Context) ([]*entity.SubDepartment, error) {
	var subdepts []*entity.SubDepartment
	err := r.db.Find(&subdepts).Error
	if err != nil {
		return nil, err
	}
	return subdepts, nil
}

func (r *subDepartmentRepo) GetByPage(ctx context.Context, page, limit uint32) ([]*entity.SubDepartment, error) {
	var subs []*entity.SubDepartment
	if int(page) <= 0 {
		page = 1
	}
	err := r.db.Offset(int(limit * (page - 1))).Limit(int(limit)).Find(&subs).Error
	if err != nil {
		return nil, err
	}
	return subs, nil
}
