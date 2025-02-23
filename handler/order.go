package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/razvanbucur360/orders-api/model"
	"github.com/razvanbucur360/orders-api/repository/order"
)

type Order struct {
	Repo order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	now := time.Now().UTC() 

	order := model.Order{
		OrderID: rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems: body.LineItems,
		CreatedAt: &now,
	}

	err := o.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("Failed to insert: " , err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshall: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Those are the orders")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order by ID")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order by ID")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete order by ID")
}
