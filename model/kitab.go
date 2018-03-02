package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

//Kitab struct kitab musnad Ahmad
type Kitab struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

//Validate method to validate fields of Kitab
func (k Kitab) Validate() error {
	return validation.ValidateStruct(&k,
		validation.Field(&k.Name, validation.Required),
	)
}

//GetListKitab return list available kitab with parameter limit and offset
func GetListKitab(start, count int) ([]Kitab, error) {
	k := Kitab{}
	kitabs := []Kitab{}
	rows, err := db.Queryx("SELECT * FROM Kitab WHERE status=1 LIMIT ? OFFSET ?", start, count)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&k)
		if err != nil {
			return nil, err
		}
		kitabs = append(kitabs, k)
	}
	return kitabs, nil
}

//Get Kitab is a method to get one kitab based on ID params
func (k *Kitab) Get() error {
	return db.QueryRow("SELECT * FROM Kitab WHERE status=1 AND id=?", k.ID).
		Scan(&k.ID, &k.Name, &k.Status, &k.CreatedAt, &k.UpdatedAt)
}

//Create new Kitab method
func (k Kitab) Create() error {
	_, err := db.NamedExec("INSERT INTO kitab (name, status, created_at, updated_at) VALUES (:name,:status,:created_at,:updated_at)",
		map[string]interface{}{
			"name":       k.Name,
			"status":     1,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})

	if err != nil {
		return err
	}
	return nil
}

//Update Kitab method
func (k Kitab) Update() error {
	_, err := db.Exec("UPDATE Kitab SET name=? WHERE id=?",
		k.Name, k.ID)

	return err
}

//Delete Kitab method
func (k Kitab) Delete() error {
	_, err := db.Exec("UPDATE Kitab SET status=? WHERE id=?",
		k.Status, k.ID)

	return err
}

//GetArchivedKitab return list available kitab with parameter limit and offset
func GetArchivedKitab(start, count int) ([]Kitab, error) {
	k := Kitab{}
	kitabs := []Kitab{}
	rows, err := db.Queryx("SELECT * FROM Kitab WHERE status=0 LIMIT ? OFFSET ?", start, count)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&k)
		if err != nil {
			return nil, err
		}
		kitabs = append(kitabs, k)
	}
	return kitabs, nil
}
