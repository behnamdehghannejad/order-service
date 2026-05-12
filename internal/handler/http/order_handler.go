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
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.CreateOrderRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.orderService.Create(req)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	order, err := h.orderService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func toCreateOrderResponse(order domain.Order) dto.CreateOrderResponse {
	return dto.CreateOrderResponse{
		UserID:    order.ID,
		Amount:    order.Amount,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
