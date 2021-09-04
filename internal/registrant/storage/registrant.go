package storage

import (
	"github.com/SemmiDev/go-pmb/internal/registrant/domain"
	"time"
)

type RegistrantResult struct {
	RegistrantId  string               `json:"registrant_id,omitempty"`
	Name          string               `json:"name,omitempty"`
	Email         string               `json:"email,omitempty"`
	Phone         string               `json:"phone,omitempty"`
	Username      string               `json:"username,omitempty"`
	Password      string               `json:"password,omitempty"`
	Code          string               `json:"code,omitempty"`
	PaymentURL    string               `json:"payment_url,omitempty"`
	Program       domain.Program       `json:"program,omitempty"`
	Bill          domain.Bill          `json:"bill,omitempty"`
	PaymentStatus domain.PaymentStatus `json:"payment_status,omitempty"`
	CreatedDate   time.Time            `json:"created_date"`
	LastUpdated   time.Time            `json:"last_updated"`
}
