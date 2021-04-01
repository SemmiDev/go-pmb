package model

type CreateProductRequest struct {
	Id			string 	`json:"id"`
	Code 		string 	`json:"code"`
	Name 		string 	`json:"name"`
	Price 		float64 `json:"price"`
	Avaliable 	bool 	`json:"avaliable"`
	Stock 		int 	`json:"stock"`
}

type CreateProductResponse struct {
	Id			string 	`json:"id"`
	Code 		string 	`json:"code"`
	Name 		string 	`json:"name"`
	Price 		float64 `json:"price"`
	Avaliable 	bool 	`json:"avaliable"`
	Stock 		int 	`json:"stock"`
}

type GetProductResponse struct {
	Id			string 	`json:"id"`
	Code 		string 	`json:"code"`
	Name 		string 	`json:"name"`
	Price 		float64 `json:"price"`
	Avaliable 	bool 	`json:"avaliable"`
	Stock 		int 	`json:"stock"`
}
