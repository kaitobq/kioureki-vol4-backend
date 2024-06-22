package models

type PlayerCategory struct {
	PlayerID uint `json:"player_id"`
	CategoryID uint `json:"category_id"`
	CategoryOptionID uint `json:"category_option_id"`
}