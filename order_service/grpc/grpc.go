package grpc

import (
	"order_service/config"
	order_service "order_service/genproto"
	"order_service/grpc/service"

	"order_service/pkg/logger"
	"order_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	order_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg))
	order_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
