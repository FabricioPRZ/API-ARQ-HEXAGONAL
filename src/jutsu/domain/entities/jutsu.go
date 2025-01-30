package entities

type Jutsu struct {
	ID              int32  `json:"id" gorm:"column:id_jutsu;primary_key"`
	Name            string `json:"name"`
	JutsuType       string `json:"jutsu_type"`
	Nature          string `json:"nature"`
	DifficultyLevel string `json:"difficulty_level"`
	CreatedBy       string `json:"created_by"`
}
