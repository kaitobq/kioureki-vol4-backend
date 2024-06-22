package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MedicalRecord struct {
	gorm.Model
	PlayerID 	uint   `json:"player_id"`
	Player 		Player `json:"player" gorm:"foreignkey:PlayerID;association_foreignkey:ID"`
	Part 		string `json:"part"`
	Diagnosis 	string `json:"diagnosis"`
	Status 		string `json:"status"`
	InjuryDate 	time.Time `json:"injury_date"`
	RecoveryDate *time.Time `json:"recovery_date"`
	Memo 		string `json:"memo"`
	OrganizationID uint   `json:"organization_id"`
	AddedBy 	string `json:"added_by"`
}
