package update

import "gorest-10/models/item/database"

type ItemUpdate struct {
	Merk  string `json:"merk" gorm:"not null"`            //varchar 100
	Jenis string `json:"jenis" gorm:"not null"`           //varchar 100
	Stok  uint8  `json:"stok" gorm:"type:uint not null"`  //tinyint unsigned
	Harga uint   `json:"harga" gorm:"type:uint not null"` //int unsigned
}

func (itemUpdate *ItemUpdate) MapFromDatabase(itemDatabase database.Item) {
	itemUpdate.Merk = itemDatabase.Merk
	itemUpdate.Jenis = itemDatabase.Jenis
	itemUpdate.Stok = itemDatabase.Stok
	itemUpdate.Harga = itemDatabase.Harga
}
