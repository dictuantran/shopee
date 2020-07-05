package models

import "github.com/dictuantran/shopee/api/config"

type OrderDetail struct {
	OrderID      int     `json:"OrderID"`
	ProductID    int     `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	ProductAlias string  `json:"ProductAlias"`
	Quantity     int     `json:"Quantity"`
	Price        float64 `json:"Price"`
}

func FetchOrderDetailByOrderID(orderID int) []OrderDetail {
	rows, err := config.DB.Query("call GetOrderDetail(?)", orderID)
	checkErr(err)

	defer rows.Close()

	ods := make([]OrderDetail, 0)
	for rows.Next() {
		od := OrderDetail{}
		err := rows.Scan(&od.OrderID,
			&od.ProductID,
			&od.ProductName,
			&od.ProductAlias,
			&od.Price,
			&od.Quantity)
		checkErr(err)

		ods = append(ods, od)
	}
	return ods
}
