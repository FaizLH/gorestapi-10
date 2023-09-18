package update

import "gorest-10/models/item/database"

type ItemUpdate struct {
	Merk  string `json:"merk" gorm:"size:100;not null"`
	Jenis string `json:"jenis" gorm:"size:100;not null"`
	Stok  uint8  `json:"stok" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Harga uint32 `json:"harga" gorm:"type:MEDIUMINT UNSIGNED NOT NULL"`
}

func (itemUpdate *ItemUpdate) MapFromDatabase(itemDatabase database.Item) {
	itemUpdate.Merk = itemDatabase.Merk
	itemUpdate.Jenis = itemDatabase.Jenis
	itemUpdate.Stok = itemDatabase.Stok
	itemUpdate.Harga = itemDatabase.Harga
}
