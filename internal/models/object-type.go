package models

// ObjectType representa um tipo de objeto
type ObjectType struct {
    TypeName string `json:"type_name" gorm:"primaryKey"`
    Name     string `json:"name"`
}