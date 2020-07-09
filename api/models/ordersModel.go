package models

import (
	"time"
)

// Order struct
type Order struct {
	OrderId       int       `json:"OrderId"`
	OrderDate     time.Time `json:"OrderDate"`
	CreatedBy     string    `json:"CreatedBy"`
	OrderStatus   string    `json:"OrderStatus"`
	PaymentMethod string    `json:"PaymentMethod"`
	PaymentStatus string    `json:"PaymentStatus"`
	Price         float64   `json:"Price"`
}

// OrderDetail struct
type OrderDetail struct {
	OrderID      int     `json:"OrderID"`
	ProductID    int     `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	ProductAlias string  `json:"ProductAlias"`
	Quantity     int     `json:"Quantity"`
	Price        float64 `json:"Price"`
}

// OrderResponse struct
type OrderResponse struct {
	RespCode string  `json:"response_code"`
	RespDesc string  `json:"response_description"`
	Data     []Order `json:"data"`
}

// OrderDetailResponse struct
type OrderDetailResponse struct {
	RespCode string        `json:"response_code"`
	RespDesc string        `json:"response_description"`
	Data     []OrderDetail `json:"data"`
}
