package models

import "github.com/dictuantran/shopee/api/config"

type Orders struct {
	StoreId        int     `json:"store_id"`
	TotalTaxAmount float64 `json:"total_tax_amount"`
	TotalAmount    float64 `json:"total_amount"`
	GrandTotal     float64 `json:"grand_total"`
}

func TotalBillByStoreIdForDraftOrder(storeId int) Orders {
	total := Orders{}

	sql := `SELECT store_id, 
		SUM(amount) as total_amount, 
		SUM(tax_amount) as total_tax_amount, 
		SUM(total_amount) as grand_total 
	FROM order_details where order_status = 0 and store_id = ? 
	GROUP BY store_id`

	row := config.DB.QueryRow(sql, storeId)

	err := row.Scan(&total.StoreId, &total.TotalAmount, &total.TotalTaxAmount, &total.GrandTotal)
	checkErr(err)

	return total
}
