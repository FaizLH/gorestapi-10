package database

import "gorest-10/models/item/request/create"

type Item struct {
	Id    uint   `json:"id" gorm:"type:MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	Merk  string `json:"merk" gorm:"size:50;not null"`
	Jenis string `json:"jenis" gorm:"size:50;not null"`
	Stok  uint8  `json:"stok" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Harga uint32 `json:"harga" gorm:"type:MEDIUMINT UNSIGNED NOT NULL"`
}

func (item *Item) MapFromRegister(itemCreate create.ItemCreate) {
	item.Merk = itemCreate.Merk
	item.Jenis = itemCreate.Jenis
	item.Stok = itemCreate.Stok
	item.Harga = itemCreate.Harga
}
