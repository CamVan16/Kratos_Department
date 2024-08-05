package service

import (
	v1 "DM/api/helloworld/v1"
	"DM/internal/biz"
	"context"
)

type EmployeeService struct {
	v1.UnimplementedEmployeeServiceServer
	biz *biz.EmployeeService
}

func NewEmployeeService(biz *biz.EmployeeService) *EmployeeService {
	return &EmployeeService{biz: biz}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, req *v1.CreateEmployeeRequest) (*v1.CreateEmployeeRespone, error) {
	emp, err := s.biz.CreateEmployee(ctx, req.Name, req.Phone, req.Password, req.Role, req.Salary, req.SubDepartmentId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEmployeeRespone{Id: emp.IDEm}, nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, req *v1.GetEmployeeByIDResquest) (*v1.GetEmployeeByIDRespone, error) {
	emp, err := s.biz.GetEnployeeByID(ctx, req.Id)
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
	err := s.biz.UpdateEmployee(ctx, req.Id, req.Name, req.Phone, req.Password, req.Role, req.Salary, req.SubDepartmentId)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateEmployeeRespone{}, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, req *v1.DeleteEmployeeResquest) (*v1.DeleteEmployeeRespone, error) {
	err := s.biz.DeleteEmployee(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteEmployeeRespone{}, nil
}

func (s *EmployeeService) GetAllEmployee(ctx context.Context, req *v1.GetAllEmployeeRequest) (*v1.GetAllEmployeeRespone, error) {
	emps, err := s.biz.GetAllEmployee(ctx)
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
	emps, err := s.biz.GetEmployeeByPhonePass(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.GetEmployeeByPhonePassRespone{
		Phone:    emps.Phone,
		Password: emps.Password,
		Role:     emps.Role,
	}, nil
}
