package entities

type Music struct {
	Name     string `gorm:"type:varchar(255); NOT NULL" json:"name"`
	Duration string `gorm:"type:varchar(6); NOT NULL" json:"duration"`
	Singer   string `gorm:"type:varchar(255); NOT NULL" json:"duration"`
	Genre    string `gorm:"type:varchar(255); NOT NULL" json:"genre"`
}
