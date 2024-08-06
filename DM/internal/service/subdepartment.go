package service

import (
	v1 "DM/api/sub_department/v1"
	"DM/internal/entity"
	"context"
)

type subdepartmentUC interface {
	CreateSubDepartment(ctx context.Context, name string, department_id uint32) (uint32, error)
	GetSubDepartmentByID(ctx context.Context, id uint32) (*entity.SubDepartment, error)
	GetAllSubDepartment(ctx context.Context) ([]*entity.SubDepartment, error)
	UpdateSubDepartment(ctx context.Context, id uint32, name string, department_id uint32) error
	DeleteSubDepartment(ctx context.Context, id uint32) error
	GetSubDepartmentByPage(ctx context.Context, page, limit uint32) ([]*entity.SubDepartment, error)
}
type SubDepartmentService struct {
	v1.UnimplementedSubDepartmentServiceServer
	//biz *biz.SubDepartmentService
	//uc *biz.SubDepartmentUC
	uc subdepartmentUC
}

func NewSubDepartmentService(uc subdepartmentUC) *SubDepartmentService {
	return &SubDepartmentService{uc: uc}
}

func (s *SubDepartmentService) CreateSubDepartment(ctx context.Context, req *v1.CreateSubDepartmentRequest) (*v1.CreateSubDepartmentRespone, error) {
	id, err := s.uc.CreateSubDepartment(ctx, req.Name, req.DepartmentId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateSubDepartmentRespone{Id: id}, nil
}

func (s *SubDepartmentService) GetSubDepartmentByID(ctx context.Context, req *v1.GetSubDepartmentByIDResquest) (*v1.GetSubDepartmentByIDRespone, error) {
	subdept, err := s.uc.GetSubDepartmentByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetSubDepartmentByIDRespone{
		Id:   subdept.ID,
		Name: subdept.Name,
	}, nil
}

func (s *SubDepartmentService) UpdateSubDepartment(ctx context.Context, req *v1.UpdateSubDepartmentResquest) (*v1.UpdateSubDepartmentRespone, error) {
	err := s.uc.UpdateSubDepartment(ctx, req.Id, req.Name, req.DepartmentId)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSubDepartmentRespone{}, nil
}

func (s *SubDepartmentService) DeleteSubDepartment(ctx context.Context, req *v1.DeleteSubDepartmentResquest) (*v1.DeleteSubDepartmentRespone, error) {
	err := s.uc.DeleteSubDepartment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSubDepartmentRespone{}, nil
}

func (s *SubDepartmentService) GetAllSubDepartment(ctx context.Context, req *v1.GetAllSubDepartmentRequest) (*v1.GetAllSubDepartmentRespone, error) {
	subdepts, err := s.uc.GetAllSubDepartment(ctx)
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

func (s *SubDepartmentService) GetSubDepartmentByPage(ctx context.Context, req *v1.GetSubDepartmentByPageRequest) (*v1.GetSubDepartmentByPageRespone, error) {
	subs, err := s.uc.GetSubDepartmentByPage(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	var subdepartments []*v1.SubDepartment
	for _, sub := range subs {
		subdepartments = append(subdepartments, &v1.SubDepartment{
			Id:           sub.ID,
			Name:         sub.Name,
			DepartmentId: sub.DepartmentID,
		})
	}

	return &v1.GetSubDepartmentByPageRespone{Subdepartments: subdepartments}, err
}
