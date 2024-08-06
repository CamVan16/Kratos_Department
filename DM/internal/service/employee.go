package service

import (
	v1 "DM/api/employee/v1"
	"DM/internal/entity"
	"context"
)

type employeeUC interface {
	CreateEmployee(ctx context.Context, name, phone, password, role string, salary uint32, sub_department_id uint32) (*entity.Employee, error)
	GetEnployeeByID(ctx context.Context, id uint32) (*entity.Employee, error)
	GetAllEmployee(ctx context.Context) ([]*entity.Employee, error)
	UpdateEmployee(ctx context.Context, id uint32, name, phone, password, role string, salary uint32, sub_department_id uint32) error
	DeleteEmployee(ctx context.Context, id uint32) error
	GetEmployeeByPhonePass(ctx context.Context, phone string) (entity.Employee, error)
	GetEmployeeByPage(ctx context.Context, page, limit uint32) ([]*entity.Employee, error)
}
type EmployeeService struct {
	v1.UnimplementedEmployeeServiceServer
	//uc *biz.EmployeeUC
	uc employeeUC
}

func NewEmployeeService(uc employeeUC) *EmployeeService {
	return &EmployeeService{uc: uc}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, req *v1.CreateEmployeeRequest) (*v1.CreateEmployeeRespone, error) {
	emp, err := s.uc.CreateEmployee(ctx, req.Name, req.Phone, req.Password, req.Role, req.Salary, req.SubDepartmentId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEmployeeRespone{Id: emp.IDEm}, nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, req *v1.GetEmployeeByIDResquest) (*v1.GetEmployeeByIDRespone, error) {
	emp, err := s.uc.GetEnployeeByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetEmployeeByIDRespone{
		Id:              emp.IDEm,
		Name:            emp.Name,
		Phone:           emp.Phone,
		Password:        emp.Password,
		Role:            emp.Role,
		Salary:          uint32(emp.Salary),
		SubDepartmentId: emp.SubDepartmentID,
	}, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, req *v1.UpdateEmployeeResquest) (*v1.UpdateEmployeeRespone, error) {
	err := s.uc.UpdateEmployee(ctx, req.Id, req.Name, req.Phone, req.Password, req.Role, req.Salary, req.SubDepartmentId)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateEmployeeRespone{}, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, req *v1.DeleteEmployeeResquest) (*v1.DeleteEmployeeRespone, error) {
	err := s.uc.DeleteEmployee(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteEmployeeRespone{}, nil
}

func (s *EmployeeService) GetAllEmployee(ctx context.Context, req *v1.GetAllEmployeeRequest) (*v1.GetAllEmployeeRespone, error) {
	emps, err := s.uc.GetAllEmployee(ctx)
	if err != nil {
		return nil, err
	}
	var employees []*v1.Employee
	for _, emps := range emps {
		employees = append(employees, &v1.Employee{
			Id:              emps.IDEm,
			Name:            emps.Name,
			Phone:           emps.Phone,
			Password:        emps.Password,
			Role:            emps.Role,
			Salary:          uint32(emps.Salary),
			SubDepartmentId: emps.SubDepartmentID,
		})
	}
	return &v1.GetAllEmployeeRespone{Employees: employees}, nil
}

func (s *EmployeeService) GetEmployeeByPhonePass(ctx context.Context, req *v1.GetEmployeeByPhonePassRequest) (*v1.GetEmployeeByPhonePassRespone, error) {
	emps, err := s.uc.GetEmployeeByPhonePass(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.GetEmployeeByPhonePassRespone{
		Phone:    emps.Phone,
		Password: emps.Password,
		Role:     emps.Role,
	}, nil
}

func (s *EmployeeService) GetEmployeeByPage(ctx context.Context, req *v1.GetEmployeeByPageRequest) (*v1.GetEmployeeByPageRespone, error) {
	emps, err := s.uc.GetEmployeeByPage(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	var employees []*v1.Employee
	for _, emp := range emps {
		employees = append(employees, &v1.Employee{
			Id:              emp.IDEm,
			Name:            emp.Name,
			Phone:           emp.Phone,
			Password:        emp.Password,
			Salary:          uint32(emp.Salary),
			Role:            emp.Role,
			SubDepartmentId: emp.SubDepartmentID,
		})
	}

	return &v1.GetEmployeeByPageRespone{Employees: employees}, err
}
