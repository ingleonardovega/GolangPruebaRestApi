package product

type Product struct {
	Id          int     `json:"id"`
	ProductCode string  `json:"ProductCode"`
	ProductName string  `json:"ProductName"`
	Description string  `json:"Description"`
	StandarCost float64 `json:"StandarCost"`
	ListPrice   float64 `json:"ListPrice"`
	Category    string  `json:"Category"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
