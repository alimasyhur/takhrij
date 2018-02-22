package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//Kitab struct
type Kitab struct {
	Table string `json:"table"`
	Name  string `json:"name"`
}

//Kutub struct list of Kitab
type Kutub struct {
	Tables []Kitab
}

func kitab(resp http.ResponseWriter, req *http.Request) {
	log.Println(getKitab())

	listKitab, _ := json.Marshal(getKitab())

	resp.Write([]byte(listKitab))
}

//getKitab funtion show list of available Kitab
func getKitab() []Kitab {

	tables := []string{}
	kutub := Kutub{}

	for _, val := range tables {
		if strings.Contains(val, "Hadits_") {
			trimmedPrefix := strings.TrimPrefix(val, "Hadits_")
			kutub.AddKitab(Kitab{
				Table: val,
				Name:  strings.Replace(trimmedPrefix, "_", " ", 1),
			})
		}
	}

	return kutub.Tables
}

//AddKitab method
func (kutub *Kutub) AddKitab(kitab Kitab) []Kitab {
	kutub.Tables = append(kutub.Tables, kitab)
	return kutub.Tables
}
