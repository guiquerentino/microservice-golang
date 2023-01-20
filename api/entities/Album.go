package entities

type Album struct {
	ID     uint   `gorm:"primaryKey; autoIncrement = true" json:"id"`
	Name   string `gorm:"type:varchar(255); NOT NULL" json:"name"`
	Genre  string `gorm:"type:varchar(255); NOT NULL" json:"genre"`
	Author string `gorm:"type:varchar(255); NOT NULL" json:"author"`
	Musics string `gorm:"type:JSON; NOT NULL" json:musics`
}

type AlbumDTO struct {
	Name   string  `gorm:"type:varchar(255); NOT NULL" json:"name"`
	Genre  string  `gorm:"type:varchar(255); NOT NULL" json:"genre"`
	Author string  `gorm:"type:varchar(255); NOT NULL" json:"author"`
	Musics []Music `gorm:"type:JSON; NOT NULL" json:musics`
}
