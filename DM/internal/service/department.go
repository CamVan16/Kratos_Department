package service

import (
	v1 "DM/api/helloworld/v1"
	"DM/internal/biz"
	"context"
)

type DepartmentService struct {
	v1.UnimplementedDepartmentServiceServer
	biz *biz.DepartmentService
}

func NewDepartmentService(biz *biz.DepartmentService) *DepartmentService {
	return &DepartmentService{biz: biz}
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, req *v1.CreateDepartmentRequest) (*v1.CreateDepartmentRespone, error) {
	id, err := s.biz.CreateDepartment(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDepartmentRespone{Id: id}, nil
}

func (s *DepartmentService) GetDepartmentByID(ctx context.Context, req *v1.GetDepartmentByIDResquest) (*v1.GetDepartmentByIDRespone, error) {
	dept, err := s.biz.GetDepartmentByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetDepartmentByIDRespone{
		Id:   dept.ID,
		Name: dept.Name,
	}, nil
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, req *v1.UpdateDepartmentResquest) (*v1.UpdateDepartmentRespone, error) {
	err := s.biz.UpdateDepartment(ctx, req.Id, req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDepartmentRespone{}, nil
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, req *v1.DeleteDepartmentResquest) (*v1.DeleteDepartmentRespone, error) {
	err := s.biz.DeleteDepartment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteDepartmentRespone{}, nil
}

func (s *DepartmentService) GetAllDepartment(ctx context.Context, req *v1.GetAllDepartmentRequest) (*v1.GetAllDepartmentRespone, error) {
	depts, err := s.biz.GetAllDepartment(ctx)
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
