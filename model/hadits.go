package model

import (
	"log"

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
	Status     int    `db:"status" json:"status"`
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

// GetListHadits return list available kitab with parameter limit and offset
// params idKitab, start, count integer type
func GetListHadits(idKitab, start, count int) ([]Hadits, error) {
	h := Hadits{}
	haditses := []Hadits{}
	rows, err := db.Queryx("SELECT * FROM hadits WHERE id_kitab=? AND status=1 ORDER BY hno ASC LIMIT ? OFFSET ?",
		idKitab,
		start,
		count,
	)

	log.Println(err)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&h)
		if err != nil {
			return nil, err
		}
		haditses = append(haditses, h)
	}
	return haditses, nil
}

//Get Hadits is a method to get one kitab based on ID params
func (h *Hadits) Get() error {
	result := db.QueryRow("SELECT * FROM hadits WHERE id=?", h.ID).
		Scan(&h.ID, &h.IDKitab, &h.Nass, &h.Terjemah, &h.NassGundul, &h.Hno, &h.Status)

	log.Println(result)

	return result
}
