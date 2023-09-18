package response

import "gorest-10/models/item/database"

type ItemResponse struct {
	Merk  string `json:"merk" gorm:"not null"`            //varchar 100
	Jenis string `json:"jenis" gorm:"not null"`           //varchar 100
	Stok  uint8  `json:"stok" gorm:"type:uint not null"`  //tinyint unsigned
	Harga uint   `json:"harga" gorm:"type:uint not null"` //int unsigned
}

func (itemResponse *ItemResponse) MapFromDatabase(itemDatabase database.Item) {
	itemResponse.Merk = itemDatabase.Merk
	itemResponse.Jenis = itemDatabase.Jenis
	itemResponse.Stok = itemDatabase.Stok
	itemResponse.Harga = itemDatabase.Harga
}
