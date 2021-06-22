package model

type (
	AuthRequest struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	CreateStudentRequest struct {
		ID                 string `json:"-"`
		FullName           string `json:"full_name"`
		Email              string `json:"email"`
		PhoneNumber        string `json:"phone_number"`
		Path               uint   `json:"path"`
		Year               uint32 `json:"year"`
		RegistrationNumber string `json:"registration_number"`
	}

	CreateStudentResponse struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	GetStudentResponse struct {
		Id                 string `json:"id"`
		Identifier         string `json:"identifier"`
		FullName           string `json:"full_name"`
		Email              string `json:"email"`
		EmailStudent       string `json:"email_student"`
		PhoneNumber        string `json:"phone_number"`
		Path               string `json:"path"`
		RegistrationNumber string `json:"registration_number"`
	}

	GetSingleStudentResponse struct {
		Id                 string `json:"id"`
		Identifier         string `json:"identifier"`
		FullName           string `json:"full_name"`
		Email              string `json:"email"`
		EmailStudent       string `json:"email_student"`
		PhoneNumber        string `json:"phone_number"`
		Path               string `json:"path"`
		RegistrationNumber string `json:"registration_number"`
		Password           string `json:"-"`
	}
)
