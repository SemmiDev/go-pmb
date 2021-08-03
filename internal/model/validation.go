package model

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

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
	if strings.HasPrefix(r.Phone, "08") == false {
		errorCollections["invalid_Phone"] = "Phone Number Is Not Valid"
	}

	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}

func (u *UpdateStatus) Validate() map[string]string {
	var errorCollections = make(map[string]string)

	if fmt.Sprint(u.VirtualAccount) == "" {
		errorCollections["Required_VA"] = "Virtual Account Is Empty"
	}
	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}

func (r *LoginRequest) Validate() map[string]string {
	var errorCollections = make(map[string]string)

	if fmt.Sprint(r.Username) == "" {
		errorCollections["Required_Username"] = "Username Is Empty"
	}

	if fmt.Sprint(r.Password) == "" {
		errorCollections["Required_Password"] = "Password Is Empty"
	}

	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}
