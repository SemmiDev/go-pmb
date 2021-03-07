package model

type CreateStudentRequest struct {
	Id string 			`json:"id"`
	Identifier string 	`json:"identifier"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
}

type CreateStudentResponse struct {
	Id string 			`json:"id"`
	Identifier string 	`json:"identifier"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
}

type GetStudentResponse struct {
	Id string 			`json:"id"`
	Identifier string 	`json:"identifier"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
}