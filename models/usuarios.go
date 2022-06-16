package models

import "time"

type Usuarios struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Folio       string    `json:"folio" gorm:"unique"`
	Nombre      string    `json:"nombre"`
	Vigencia    time.Time `json:"vigencia"`
	Areas       string    `json:"areas"`
	Privilegios string    `json:"privilegios"`
}
