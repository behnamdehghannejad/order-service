package http

import (
	"encoding/json"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/handler/dto"
	"order-service/internal/service"
	"strconv"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (handler *OrderHandler) Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var req dto.CreateOrderRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Create(toDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer)
}

func (handler *OrderHandler) GetById(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	order, err := handler.service.GetByID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	returnResponse(writer, order)
}

func (handler *OrderHandler) GetByUserId(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.PathValue("user-id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
	}

	order, err := handler.service.GetByUserId(userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
	}

	returnResponse(writer, order)
}

func (handler *OrderHandler) listAll(writer http.ResponseWriter, request *http.Request) {
	all, err := handler.service.ListAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(toResponseList(all)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *OrderHandler) updateStatus(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
	}

	var req dto.OrderStatusRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	order, err := handler.service.UpdateStatus(id, domain.Status(req.Status))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	returnResponse(writer, order)
}

func returnResponse(writer http.ResponseWriter, order domain.Order) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(toOrderResponse(order)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *OrderHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
	}

	if err := handler.service.Delete(id); err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
	}

	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer)
}

func toOrderResponse(order domain.Order) dto.OrderResponse {
	return dto.OrderResponse{
		UserID:    order.ID,
		Amount:    order.Amount,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}

func toDomain(order dto.CreateOrderRequest) domain.Order {
	return domain.Order{
		UserID: order.UserID,
		Amount: order.Amount,
		Status: domain.Status(order.Status),
	}
}

func toResponseList(all []domain.Order) []dto.OrderResponse {
	allOrderResponse := make([]dto.OrderResponse, 0, len(all))
	for i, order := range all {
		allOrderResponse[i] = toOrderResponse(order)
	}
	return allOrderResponse
}
