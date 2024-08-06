package service

import (
	v1 "DM/api/department/v1"
	"DM/internal/entity"
	"context"
)

type departmentUC interface {
	CreateDepartment(ctx context.Context, name string) (uint32, error)
	GetDepartmentByID(ctx context.Context, id uint32) (*entity.Department, error)
	UpdateDepartment(ctx context.Context, id uint32, name string) error
	DeleteDepartment(ctx context.Context, id uint32) error
	GetAllDepartment(ctx context.Context) ([]*entity.Department, error)
	GetDepartmentByPage(ctx context.Context, page, limit uint32) ([]*entity.Department, error)
}

type DepartmentService struct {
	v1.UnimplementedDepartmentServiceServer
	//uc *biz.DepartmentUC
	uc departmentUC
}

func NewDepartmentService(uc departmentUC) *DepartmentService {
	return &DepartmentService{uc: uc}
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, req *v1.CreateDepartmentRequest) (*v1.CreateDepartmentRespone, error) {
	id, err := s.uc.CreateDepartment(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDepartmentRespone{Id: id}, nil
}

func (s *DepartmentService) GetDepartmentByID(ctx context.Context, req *v1.GetDepartmentByIDResquest) (*v1.GetDepartmentByIDRespone, error) {
	dept, err := s.uc.GetDepartmentByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetDepartmentByIDRespone{
		Id:   dept.ID,
		Name: dept.Name,
	}, nil
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, req *v1.UpdateDepartmentResquest) (*v1.UpdateDepartmentRespone, error) {
	err := s.uc.UpdateDepartment(ctx, req.Id, req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDepartmentRespone{}, nil
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, req *v1.DeleteDepartmentResquest) (*v1.DeleteDepartmentRespone, error) {
	err := s.uc.DeleteDepartment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteDepartmentRespone{}, nil
}

func (s *DepartmentService) GetAllDepartment(ctx context.Context, req *v1.GetAllDepartmentRequest) (*v1.GetAllDepartmentRespone, error) {
	depts, err := s.uc.GetAllDepartment(ctx)
	if err != nil {
		return nil, err
	}
	var departments []*v1.Department
	for _, dept := range depts {
		departments = append(departments, &v1.Department{
			Id:   dept.ID,
			Name: dept.Name,
		})
	}
	return &v1.GetAllDepartmentRespone{Departments: departments}, nil
}

func (s *DepartmentService) GetDepartmentByPage(ctx context.Context, req *v1.GetDepartmentByPageRequest) (*v1.GetDepartmentByPageRespone, error) {
	depts, err := s.uc.GetDepartmentByPage(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	var departments []*v1.Department
	for _, dept := range depts {
		departments = append(departments, &v1.Department{
			Id:   dept.ID,
			Name: dept.Name,
		})
	}
	return &v1.GetDepartmentByPageRespone{Departments: departments}, nil
}
