package main

import (
	"net/http"

	common "github.com/arturfil/m_commons"
	pb "github.com/arturfil/m_commons/api"
)

type handler struct {
    // gateway
    client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
    return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /api/v1/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
    customerID := r.PathValue("customerID")

    var items []*pb.ItemsWithQuantity
    if err := common.ReadJSON(r, &items); err != nil {
        common.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
        CustomerID: customerID,
        Items: items,
    })
}
