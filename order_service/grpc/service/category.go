package service

import (
	"context"
	"order_service/config"
	order_service "order_service/genproto"
	"order_service/pkg/logger"
	"order_service/storage"
)

type CategoryService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	order_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *CategoryService {
	return &CategoryService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *CategoryService) Create(ctx context.Context, req *order_service.CreateCategoryRequest) (*order_service.Response, error) {
	id, err := b.storage.Category().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: id}, nil
}

func (b *CategoryService) Get(ctx context.Context, req *order_service.IdRequest) (*order_service.Category, error) {
	reso, err := b.storage.Category().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *CategoryService) List(ctx context.Context, req *order_service.ListCategoryRequest) (*order_service.ListCategoryResponse, error) {
	Categorys, err := b.storage.Category().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.ListCategoryResponse{Categories: Categorys.Categories,
		Count: Categorys.Count}, nil
}

func (s *CategoryService) Update(ctx context.Context, req *order_service.UpdateCategoryRequest) (*order_service.Response, error) {
	resp, err := s.storage.Category().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}

func (s *CategoryService) Delete(ctx context.Context, req *order_service.IdRequest) (*order_service.Response, error) {
	resp, err := s.storage.Category().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}
