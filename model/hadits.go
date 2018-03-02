package model

import (
	"github.com/go-ozzo/ozzo-validation"
)

//Hadits struct kitab musnad Ahmad
type Hadits struct {
	ID         int    `db:"id" json:"id"`
	IDKitab    int    `db:"id_kitab" json:"id_kitab"`
	Nass       string `db:"nass" json:"nass"`
	Terjemah   string `db:"terjemah" json:"terjemah"`
	NassGundul string `db:"nass_gundul" json:"nassgundul"`
	Hno        int    `db:"hno" json:"hno"`
}

//Validate method to validate fields of Hadits
func (h Hadits) Validate() error {
	return validation.ValidateStruct(&h,
		validation.Field(&h.Nass, validation.Required),
		validation.Field(&h.Terjemah, validation.Required),
		validation.Field(&h.NassGundul, validation.Required),
		validation.Field(&h.Hno, validation.Required),
	)
}
