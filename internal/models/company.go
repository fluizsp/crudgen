package models

import (
	"github.com/google/uuid"
)

// Company representa um tipo de objeto "Company"
type Company struct {
	ID              string `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	TradeName       string `json:"trade_name"`
	CNPJ            string `json:"cnpj"`
	FocalPointName  string `json:"focal_point_name"`
	FocalPointEmail string `json:"focal_point_email"`
	AddressLine1    string `json:"address_line1"`
	AddressLine2    string `json:"address_line2"`
	Complement      string `json:"complement"`
	ZipCode         string `json:"zip_code"`
	City            string `json:"city"`
	State           string `json:"state"`
}

// NewCompany cria uma nova instância de Company com um ID GUID
func NewCompany() *Company {
	return &Company{ID: uuid.New().String()}
}
