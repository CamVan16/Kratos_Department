package data

import (
	"DM/internal/biz"
	"DM/internal/models"
	"context"

	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *Data) biz.EmployeeRepo {
	return &employeeRepo{db: db.DB}
}

func (r *employeeRepo) Create(ctx context.Context, emp *models.Employee) error {
	return r.db.Create(emp).Error
}

func (r *employeeRepo) GetByID(ctx context.Context, id uint32) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.First(&emp, "id_em = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepo) Update(ctx context.Context, emp *models.Employee) error {
	return r.db.Save(emp).Error
}

func (r *employeeRepo) Delete(ctx context.Context, id uint32) error {
	return r.db.Delete(&models.Employee{}, id).Error
}

func (r *employeeRepo) GetAll(ctx context.Context) ([]*models.Employee, error) {
	var emps []*models.Employee
	err := r.db.Find(&emps).Error
	if err != nil {
		return nil, err
	}
	return emps, nil
}

func (r *employeeRepo) FindByPhone(ctx context.Context, phone string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("phone = ?", phone).First(&employee).Error
	return employee, err
}
