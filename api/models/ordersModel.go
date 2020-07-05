package models

import (
	"time"

	"github.com/dictuantran/shopee/api/config"
)

type Order struct {
	OrderId       int       `json:"OrderId"`
	OrderDate     time.Time `json:"OrderDate"`
	CreatedBy     string    `json:"CreatedBy"`
	OrderStatus   string    `json:"OrderStatus"`
	PaymentMethod string    `json:"PaymentMethod"`
	PaymentStatus string    `json:"PaymentStatus"`
	Price         float64   `json:"Price"`
}

func FetchOrder() []Order {
	rows, err := config.DB.Query("call GetOrders('2017-02-01', '2017-05-01')")
	checkErr(err)
	defer rows.Close()

	orders := make([]Order, 0)
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.OrderId,
			&order.OrderDate,
			&order.CreatedBy,
			&order.PaymentMethod,
			&order.OrderStatus,
			&order.PaymentStatus,
			&order.Price)
		checkErr(err)

		orders = append(orders, order)
	}

	return orders
}
