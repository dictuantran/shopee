package repositories

import (
	"github.com/dictuantran/shopee/api/config"
	"github.com/dictuantran/shopee/api/models"
)

// OrderRepository struct
type (
	OrderRepository struct{}
)

// FetchOrder method
func (r *OrderRepository) FetchOrder() ([]models.Order, error) {
	rows, err := config.DB.Query("call GetOrders('2017-02-01', '2017-05-01')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderId,
			&order.OrderDate,
			&order.CreatedBy,
			&order.PaymentMethod,
			&order.OrderStatus,
			&order.PaymentStatus,
			&order.Price)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// FetchOrderDetailByOrderID method
func (r *OrderRepository) FetchOrderDetailByOrderID(orderID int) ([]models.OrderDetail, error) {
	rows, err := config.DB.Query("call GetOrderDetail(?)", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ods := make([]models.OrderDetail, 0)
	for rows.Next() {
		od := models.OrderDetail{}
		err := rows.Scan(&od.OrderID,
			&od.ProductID,
			&od.ProductName,
			&od.ProductAlias,
			&od.Price,
			&od.Quantity)
		if err != nil {
			return nil, err
		}

		ods = append(ods, od)
	}
	return ods, nil
}
