package model

import (
	"html"
	"regexp"
	"strings"
)

type RegistrationRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Program string `json:"program"`
}

type RegistrationResponse struct {
	Username      string        `json:"username"`
	Password      string        `json:"password"`
	Bill          Bill          `json:"bill"`
	AccountNumber AccountNumber `json:"account_number"`
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	phoneRegexp = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
)

func (r *RegistrationRequest) Validate() map[string]string {
	var errorCollections = make(map[string]string)

	r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	r.Email = html.EscapeString(strings.TrimSpace(r.Email))
	r.Phone = html.EscapeString(strings.TrimSpace(r.Phone))

	if r.Name == "" {
		errorCollections["Required_Name"] = "Name Is Empty"
	}
	if len(r.Name) > 100 {
		errorCollections["Length_Name_Too_Long"] = "Length_Name_Too_Long"
	}
	if r.Email == "" {
		errorCollections["Required_Email"] = "Email Is Empty"
	}
	if r.Phone == "" {
		errorCollections["Required_Phone"] = "Phone Is Empty"
	}
	if r.Program == "" {
		errorCollections["Required_Program"] = "Program Is Empty"
	}

	if emailRegexp.MatchString(r.Email) == false {
		errorCollections["invalid_Email"] = "Email Is Not Valid"
	}
	if phoneRegexp.MatchString(r.Phone) == false {
		errorCollections["invalid_Phone"] = "Phone Number Is Not Valid"
	}

	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}
