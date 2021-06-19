package entity

type Student struct {
	Id                 string `bson:"id"`
	Identifier         string `bson:"identifier"`
	FullName           string `bson:"full_name"`
	Email              string `bson:"email"`
	EmailStudent       string `bson:"email_student"`
	PhoneNumber        string `bson:"phone_number"`
	Path               string `bson:"path"`
	RegistrationNumber string `bson:"registration_number"`
	Password           string `bson:"password"`
}