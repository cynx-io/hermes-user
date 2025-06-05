package entity

type TblUser struct {
	EssentialEntity
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Coin     int    `gorm:"default:0" json:"coin"`
}
