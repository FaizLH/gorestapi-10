package create

type ItemCreate struct {
	Merk  string `json:"merk" gorm:"not null"`            //varchar 100
	Jenis string `json:"jenis" gorm:"not null"`           //varchar 100
	Stok  uint8  `json:"stok" gorm:"type:uint not null"`  //tinyint unsigned
	Harga uint   `json:"harga" gorm:"type:uint not null"` //int unsigned
}
