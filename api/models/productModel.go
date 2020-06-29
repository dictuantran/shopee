package models

import "github.com/dictuantran/shopee/api/config"

// Product detail is
type Product struct {
	ID             int     `json:ID`
	Name           string  `json:Name`
	Alias          string  `json:Alias`
	CategoryID     int     `json:CategoryID`
	Image          string  `json:Image`
	Price          float64 `json:Price`
	OriginalPrice  float64 `json:OriginalPrice`
	PromotionPrice float64 `json:PromotionPrice`
	Warranty       int     `json:Warranty`
	Description    string  `json:Description`
	Content        string  `json:Content`
}

func FetchProduct() []Product {
	rows, err := config.DB.Query("call GetProducts(?)", 1)
	checkErr(err)

	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Alias, &product.CategoryID)
		checkErr(err)

		products = append(products, product)
	}
	return products
}
