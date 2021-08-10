package models

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

var (
	errorCollections = make(map[string]string)
	emailRegexp      = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	phoneRegexp      = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
)

func (r *RegistrationRequest) Validate() map[string]string {
	errorCollections = make(map[string]string)

	r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	r.Email = html.EscapeString(strings.TrimSpace(r.Email))
	r.Phone = html.EscapeString(strings.TrimSpace(r.Phone))

	if r.Name == "" {
		errorCollections["Required_Name"] = "Name Is Empty"
	}
	if len(r.Name) > 100 {
		errorCollections["Name_Too_Long"] = "Name Too Long"
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
		errorCollections["Invalid_Email"] = "Email Is Not Valid"
	}
	if phoneRegexp.MatchString(r.Phone) == false {
		errorCollections["Invalid_Phone"] = "Phone Number Is Not Valid"
	}
	if strings.HasPrefix(r.Phone, "08") == false {
		errorCollections["Invalid_Phone"] = "Phone Number Is Not Valid"
	}

	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}

func (u *UpdatePaymentStatusRequest) Validate() map[string]string {
	errorCollections = make(map[string]string)

	if u.RegisterID == "" {
		errorCollections["Required_ID"] = "ID Is Empty"
	}
	if u.PaymentStatus == "" {
		errorCollections["Required_PaymentStatus"] = "PaymentStatus Is Empty"
	}
	if u.FraudStatus == "" {
		errorCollections["Required_FraudStatus"] = "FraudStatus Is Empty"
	}
	if u.PaymentType == "" {
		errorCollections["Required_PaymentType"] = "PaymentType Is Empty"
	}
	if len(errorCollections) > 0 {
		return errorCollections
	}
	return nil
}

func (r *LoginRequest) Validate() map[string]string {
	errorCollections = make(map[string]string)

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
