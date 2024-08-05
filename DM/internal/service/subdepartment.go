package service

import (
	v1 "DM/api/helloworld/v1"
	"DM/internal/biz"
	"context"
)

type SubDepartmentService struct {
	v1.UnimplementedSubDepartmentServiceServer
	biz *biz.SubDepartmentService
}

func NewSubDepartmentService(biz *biz.SubDepartmentService) *SubDepartmentService {
	return &SubDepartmentService{biz: biz}
}

func (s *SubDepartmentService) CreateSubDepartment(ctx context.Context, req *v1.CreateSubDepartmentRequest) (*v1.CreateSubDepartmentRespone, error) {
	id, err := s.biz.CreateSubDepartment(ctx, req.Name, req.DepartmentId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateSubDepartmentRespone{Id: id}, nil
}

func (s *SubDepartmentService) GetSubDepartmentByID(ctx context.Context, req *v1.GetSubDepartmentByIDResquest) (*v1.GetSubDepartmentByIDRespone, error) {
	subdept, err := s.biz.GetSubDepartmentByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetSubDepartmentByIDRespone{
		Id:   subdept.ID,
		Name: subdept.Name,
	}, nil
}

func (s *SubDepartmentService) UpdateSubDepartment(ctx context.Context, req *v1.UpdateSubDepartmentResquest) (*v1.UpdateSubDepartmentRespone, error) {
	err := s.biz.UpdateSubDepartment(ctx, req.Id, req.Name, req.DepartmentId)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSubDepartmentRespone{}, nil
}

func (s *SubDepartmentService) DeleteSubDepartment(ctx context.Context, req *v1.DeleteSubDepartmentResquest) (*v1.DeleteSubDepartmentRespone, error) {
	err := s.biz.DeleteSubDepartment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSubDepartmentRespone{}, nil
}

func (s *SubDepartmentService) GetAllSubDepartment(ctx context.Context, req *v1.GetAllSubDepartmentRequest) (*v1.GetAllSubDepartmentRespone, error) {
	subdepts, err := s.biz.GetAllSubDepartment(ctx)
	if err != nil {
		return nil, err
	}
	var subdepartments []*v1.SubDepartment
	for _, subdepts := range subdepts {
		subdepartments = append(subdepartments, &v1.SubDepartment{
			Id:           subdepts.ID,
			Name:         subdepts.Name,
			DepartmentId: subdepts.DepartmentID,
		})
	}
	return &v1.GetAllSubDepartmentRespone{Subdepartments: subdepartments}, nil
}
