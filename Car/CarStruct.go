package Car

type Car struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:50" json:"nama"`
	Photo  string `gorm:"size:255" json:"foto"`
	Engine string `gorm:"size:10" json:"mesin"`
	Price  int    `gorm:"size:20" json:"harga"`
	Stock  int    `gorm:"size:3" json:"jumlah"`
}
type Add struct {
	Name   string `gorm:"size:50" json:"nama"`
	Photo  string `gorm:"size:255" json:"foto"`
	Engine string `gorm:"size:10" json:"mesin"`
	Price  int    `gorm:"size:20" json:"harga"`
	Stock  int    `gorm:"size:3" json:"jumlah"`
}
