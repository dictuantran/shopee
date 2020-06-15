package models

// Item is
type Item struct {
	ItemID      int     `json:id`
	Link        string  `json:link`
	Name        string  `json:name`
	Summary     string  `json:summary`
	Year        int     `json:year`
	Country     string  `json:country`
	Price       float64 `json:price`
	Description string  `json:desciption`
}
