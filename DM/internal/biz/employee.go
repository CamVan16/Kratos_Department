package biz

import (
	"DM/internal/models"
	"context"
)

type EmployeeRepo interface {
	Create(ctx context.Context, emp *models.Employee) error
	GetByID(ctx context.Context, id uint32) (*models.Employee, error)
	Update(ctx context.Context, emp *models.Employee) error
	Delete(ctx context.Context, id uint32) error
	GetAll(ctx context.Context) ([]*models.Employee, error)
	FindByPhone(ctx context.Context, phone string) (models.Employee, error)
}

type EmployeeService struct {
	repo EmployeeRepo
}

func NewEmployeeService(repo EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, name, phone, password, role string, salary uint32, sub_department_id uint32) (*models.Employee, error) {
	emp := &models.Employee{
		Name:            name,
		Phone:           phone,
		Password:        password,
		Role:            role,
		Salary:          float64(salary),
		SubDepartmentID: sub_department_id,
	}
	if err := s.repo.Create(ctx, emp); err != nil {
		return nil, err
	}
	return emp, nil
}

func (s *EmployeeService) GetEnployeeByID(ctx context.Context, id uint32) (*models.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) GetAllEmployee(ctx context.Context) ([]*models.Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, id uint32, name, phone, password, role string, salary uint32, sub_department_id uint32) error {
	emp := &models.Employee{
		IDEm:            id,
		Name:            name,
		Phone:           phone,
		Password:        password,
		Salary:          float64(salary),
		Role:            role,
		SubDepartmentID: sub_department_id,
	}
	return s.repo.Update(ctx, emp)
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}

//	func (s *employeeService) GetEmployeeByPhonePass(phone, pass string) (models.Employee, error) {
//		return s.repository.FindByPhonePass(phone, pass)
//	}
func (s *EmployeeService) GetEmployeeByPhonePass(ctx context.Context, phone string) (models.Employee, error) {
	return s.repo.FindByPhone(ctx, phone)
}
