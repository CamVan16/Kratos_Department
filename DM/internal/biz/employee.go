package biz

import (
	"DM/internal/entity"
	"context"
)

type EmployeeRepo interface {
	Create(ctx context.Context, emp *entity.Employee) error
	GetByID(ctx context.Context, id uint32) (*entity.Employee, error)
	Update(ctx context.Context, emp *entity.Employee) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*entity.Employee, error)
	FindByPhone(ctx context.Context, phone string) (entity.Employee, error)
	GetByPage(ctx context.Context, page, limit uint32) ([]*entity.Employee, error)
}

type EmployeeUC struct {
	repo EmployeeRepo
}

func NewEmployeeUC(repo EmployeeRepo) *EmployeeUC {
	return &EmployeeUC{repo: repo}
}

func (uc *EmployeeUC) CreateEmployee(ctx context.Context, name, phone, password, role string, salary uint32, sub_department_id uint32) (*entity.Employee, error) {
	emp := &entity.Employee{
		Name:            name,
		Phone:           phone,
		Password:        password,
		Role:            role,
		Salary:          float64(salary),
		SubDepartmentID: sub_department_id,
	}
	if err := uc.repo.Create(ctx, emp); err != nil {
		return nil, err
	}
	return emp, nil
}

func (uc *EmployeeUC) GetEnployeeByID(ctx context.Context, id uint32) (*entity.Employee, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *EmployeeUC) GetAllEmployee(ctx context.Context) ([]*entity.Employee, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *EmployeeUC) UpdateEmployee(ctx context.Context, id uint32, name, phone, password, role string, salary uint32, sub_department_id uint32) error {
	emp := &entity.Employee{
		IDEm:            id,
		Name:            name,
		Phone:           phone,
		Password:        password,
		Salary:          float64(salary),
		Role:            role,
		SubDepartmentID: sub_department_id,
	}
	return uc.repo.Update(ctx, emp)
}

func (uc *EmployeeUC) DeleteEmployee(ctx context.Context, id uint32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *EmployeeUC) GetEmployeeByPhonePass(ctx context.Context, phone string) (entity.Employee, error) {
	return uc.repo.FindByPhone(ctx, phone)
}

func (uc *EmployeeUC) GetEmployeeByPage(ctx context.Context, page, limit uint32) ([]*entity.Employee, error) {
	return uc.repo.GetByPage(ctx, page, limit)
}
