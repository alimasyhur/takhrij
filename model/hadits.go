package model

//MusnadAhmad struct kitab musnad Ahmad
type MusnadAhmad struct {
	ID         int    `db:"id"`
	Nass       string `db:"nass"`
	Terjemah   string `db:"terjemah"`
	NassGundul string `db:"nass_gundul"`
	Hno        int    `db:"hno"`
}
