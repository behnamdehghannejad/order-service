package handler

import (
	"context"
	"order-service/internal/domain"
	"time"

	"order-service/internal/service"
	pb "order-service/proto/generate"

	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderGrpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderGrpcHandler(service service.OrderService) *OrderGrpcHandler {
	return &OrderGrpcHandler{
		service: service,
	}
}

func (handler *OrderGrpcHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {

	requestDto := toDomain(request)

	if err := handler.service.Create(requestDto); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Order: &pb.Order{},
	}, nil
}

func (handler *OrderGrpcHandler) GetByID(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {

	order, err := handler.service.GetByID(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetByIdResponse{Order: toProto(order)}, nil
}

func (handler *OrderGrpcHandler) GetByUerID(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {

	order, err := handler.service.GetByUserId(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetByUserIdResponse{Order: toProto(order)}, nil
}

func (handler *OrderGrpcHandler) ListAll(ctx context.Context, request *emptypb.Empty) (*pb.ListOrderResponse, error) {

	orders, err := handler.service.ListAll()
	if err != nil {
		return nil, err
	}

	var result []*pb.Order
	for _, order := range orders {
		result = append(result, toProto(order))
	}

	return &pb.ListOrderResponse{Orders: result}, nil
}

func (handler *OrderGrpcHandler) UpdateStatus(ctx context.Context, request *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {

	order, err := handler.service.UpdateStatus(int(request.Id), domain.Status(request.Status))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateStatusResponse{Order: toProto(order)}, nil
}

func (handler *OrderGrpcHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*emptypb.Empty, error) {
	if err := handler.service.Delete(int(request.Id)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func toDomain(request *pb.CreateRequest) domain.Order {
	return domain.Order{
		UserID: int(request.UserId),
		Amount: request.Amount,
		Status: domain.Status(request.Status),
	}
}

func toProto(order domain.Order) *pb.Order {
	return &pb.Order{
		Id:        int32(order.ID),
		UserId:    int32(order.UserID),
		Amount:    order.Amount,
		Status:    pb.Status(pb.Status_value[string(order.Status)]),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
