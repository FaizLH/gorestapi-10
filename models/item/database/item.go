package database

import "gorest-10/models/item/request/create"

type Item struct {
	Id    uint   `json:"id" gorm:"primaryKey autoIncrement type:uint"` //int unsigned
	Merk  string `json:"merk" gorm:"not null"`                         //varchar 100
	Jenis string `json:"jenis" gorm:"not null"`                        //varchar 100
	Stok  uint8  `json:"stok" gorm:"type:uint not null"`               //tinyint unsigned
	Harga uint   `json:"harga" gorm:"type:uint not null"`              //int unsigned
}

func (item *Item) MapFromRegister(itemCreate create.ItemCreate) {
	item.Merk = itemCreate.Merk
	item.Jenis = itemCreate.Jenis
	item.Stok = itemCreate.Stok
	item.Harga = itemCreate.Harga
}
