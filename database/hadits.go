package database

//GetHaditsByIDKitab query to get on available hadits
// func GetHaditsByIDKitab(idKitab string) (hadits []model.Hadits) {
// 	data := []model.Hadits{}
// 	err := db.Select(&data, "SELECT * FROM hadits WHERE id_kitab=? LIMIT 20", idKitab)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return data
// }

// //GetOneHadits query to get on available hadits
// func GetOneHadits(id, idKitab string) (hadits model.Hadits) {
// 	data := model.Hadits{}
// 	err := db.Get(&data, "SELECT * FROM hadits WHERE id=? AND id_kitab=?", id, idKitab)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return data
// }
