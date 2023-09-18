package response

import "gorest-10/models/item/database"

type ItemResponse struct {
	Merk  string `json:"merk" gorm:"size:50;not null"`
	Jenis string `json:"jenis" gorm:"size:50;not null"`
	Stok  uint8  `json:"stok" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Harga uint32 `json:"harga" gorm:"type:MEDIUMINT UNSIGNED NOT NULL"`
}

func (itemResponse *ItemResponse) MapFromDatabase(itemDatabase database.Item) {
	itemResponse.Merk = itemDatabase.Merk
	itemResponse.Jenis = itemDatabase.Jenis
	itemResponse.Stok = itemDatabase.Stok
	itemResponse.Harga = itemDatabase.Harga
}
