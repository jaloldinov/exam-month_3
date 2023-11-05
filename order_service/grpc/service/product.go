package service

import (
	"context"
	"order_service/config"
	order_service "order_service/genproto"
	"order_service/pkg/logger"
	"order_service/storage"
)

type ProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	order_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *ProductService {
	return &ProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *ProductService) Create(ctx context.Context, req *order_service.CreateProductRequest) (*order_service.Response, error) {
	id, err := b.storage.Product().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: id}, nil
}

func (b *ProductService) Get(ctx context.Context, req *order_service.IdRequest) (*order_service.Product, error) {
	reso, err := b.storage.Product().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *ProductService) List(ctx context.Context, req *order_service.ListProductRequest) (*order_service.ListProductResponse, error) {
	Products, err := b.storage.Product().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.ListProductResponse{Products: Products.Products,
		Count: Products.Count}, nil
}

func (s *ProductService) Update(ctx context.Context, req *order_service.UpdateProductRequest) (*order_service.Response, error) {
	resp, err := s.storage.Product().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}

func (s *ProductService) Delete(ctx context.Context, req *order_service.IdRequest) (*order_service.Response, error) {
	resp, err := s.storage.Product().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}
