package models

import "github.com/jinzhu/gorm"

type Usuarios struct {
	gorm.Model
	Folio       string `json:"folio" gorm:"unique"`
	Nombre      string `json:"nombre"`
	Areas       string `json:"areas"`
	Privilegios string `json:"privilegios"`
}

//Vigencia    time.Time `json:"vigencia"`
