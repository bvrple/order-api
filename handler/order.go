package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get order")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get order")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update order")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete order")
}
