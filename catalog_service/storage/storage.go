package storage

import (
	pb "catalog_service/genproto"
	"context"
)

type StorageI interface {
	Category() CategoryI
}

type CategoryI interface {
	Create(context.Context, *pb.CreateCategoryRequest) (string, error)
	Get(context.Context, *pb.IdRequest) (*pb.Category, error)
	GetList(context.Context, *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error)
	Update(context.Context, *pb.UpdateCategoryRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)
}
