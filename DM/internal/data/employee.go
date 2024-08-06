package data

import (
	"DM/internal/biz"
	"DM/internal/entity"
	"context"

	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *Data) biz.EmployeeRepo {
	return &employeeRepo{db: db.DB}
}

func (r *employeeRepo) Create(ctx context.Context, emp *entity.Employee) error {
	return r.db.Create(emp).Error
}

func (r *employeeRepo) GetByID(ctx context.Context, id uint32) (*entity.Employee, error) {
	var emp entity.Employee
	err := r.db.First(&emp, "id_em = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepo) Update(ctx context.Context, emp *entity.Employee) error {
	return r.db.Save(emp).Error
}

func (r *employeeRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&entity.Employee{}, id).Error
}

func (r *employeeRepo) GetAll(ctx context.Context) ([]*entity.Employee, error) {
	var emps []*entity.Employee
	err := r.db.Find(&emps).Error
	if err != nil {
		return nil, err
	}
	return emps, nil
}

func (r *employeeRepo) FindByPhone(ctx context.Context, phone string) (entity.Employee, error) {
	var employee entity.Employee
	err := r.db.Where("phone = ?", phone).First(&employee).Error
	return employee, err
}

func (r *employeeRepo) GetByPage(ctx context.Context, page, limit uint32) ([]*entity.Employee, error) {
	var emps []*entity.Employee
	if page <= 0 {
		page = 1
	}
	err := r.db.Offset(int(limit * (page - 1))).Limit(int(limit)).Find(&emps).Error
	if err != nil {
		return nil, err
	}
	return emps, nil
}
