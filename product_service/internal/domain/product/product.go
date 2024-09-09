package product

type InputResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Quantity    string `json:"quantity"`
	AddedDate   string `json:"added_date"`
}

type Entity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Quantity    string `json:"quantity"`
	AddedDate   string `json:"added_date"`
}
