package entities

type Shinobi struct {
	ID              int32  `json:"id" gorm:"column:id_shinobi;primary_key"`
	Name            string `json:"name"`
	Clan            string `json:"clan"`
	Position        string `json:"position"`
	Village         string `json:"village"`
	SpecialHability string `json:"special_hability"`
}
