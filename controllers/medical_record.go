package controllers

import (
	"backend/models"
	"backend/utils/token"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateMedicalRecordInput struct {
	PlayerID 		uint   `json:"player_id" binding:"required"`
	Part 			string `json:"part" binding:"required"`
	Diagnosis 		string `json:"diagnosis"`
	Status 			string `json:"status"`
	InjuryDate      string `json:"injury_date" binding:"required"`
	RecoveryDate 	string `json:"recovery_date"`
	Memo 			string `json:"memo"`
}

func CreateMedicalRecord(c *gin.Context){
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	// AddedByをUserから探す
	var user models.User
	err = models.DB.Find(&user, userId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input CreateMedicalRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02"
	injuryDate, err := time.Parse(layout, input.InjuryDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recoveryDate *time.Time
	if input.RecoveryDate != "" {
		parsedDate, err := time.Parse(layout, input.RecoveryDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		recoveryDate = &parsedDate
	}

	// 負傷した選手を探す
	var player models.Player
	err = models.DB.Find(&player, input.PlayerID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	organizationId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var medicalRecord models.MedicalRecord
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 記録を作成
		medicalRecord = models.MedicalRecord{
			PlayerID: 	 	input.PlayerID,
			Player:      	player,
			Part:           input.Part,
			Diagnosis:      input.Diagnosis,
			Status:         input.Status,
			InjuryDate:     injuryDate,
			RecoveryDate:   recoveryDate,
			Memo:           input.Memo,
			OrganizationID: uint(organizationId),
			AddedBy:        user.Username,
		}
		if err := tx.Create(&medicalRecord).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"added_record": medicalRecord})
}

type UpdateMedicalRecordInput struct {
	Part 			string `json:"part" binding:"required"`
	Diagnosis 		string `json:"diagnosis"`
	Status 			string `json:"status"`
	InjuryDate      string `json:"injury_date" binding:"required"`
	RecoveryDate 	string `json:"recovery_date"`
	Memo 			string `json:"memo"`
}

func UpdateMedicalRecord(c *gin.Context){
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	organizationId := c.Param("id")
	medicalRecordId := c.Param("record_id")
	var medicalRecord models.MedicalRecord

	if err := models.DB.Where("id = ? AND organization_id = ?", medicalRecordId, organizationId).Preload("Player").First(&medicalRecord).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input UpdateMedicalRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02"
	injuryDate, err := time.Parse(layout, input.InjuryDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recoveryDate *time.Time
	if input.RecoveryDate != "" {
		parsedDate, err := time.Parse(layout, input.RecoveryDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		recoveryDate = &parsedDate
	}

	var user models.User
	err = models.DB.Find(&user, userId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.DB.Model(&medicalRecord).Updates(models.MedicalRecord{
		Part:           input.Part,
		Diagnosis:      input.Diagnosis,
		Status:         input.Status,
		InjuryDate:     injuryDate,
		RecoveryDate:   recoveryDate,
		Memo:           input.Memo,
		AddedBy:        user.Username,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_record": medicalRecord})
}

func GetMedicalRecords(c *gin.Context) {
	organizationId := c.Param("id")
	var medicalRecords []models.MedicalRecord
	if err := models.DB.Where("organization_id = ?", organizationId).Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

func GetInjuredToday(c *gin.Context){
	organizationId := c.Param("id")
	var medicalRecords []models.MedicalRecord
	today := time.Now().Format("2006-01-02")

	startOfDay, err := time.Parse("2006-01-02", today)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	if err := models.DB.Where("organization_id = ? AND injury_date >= ? AND injury_date < ?", organizationId, startOfDay, endOfDay).Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

// 直近1週間の負傷者を取得
func GetInjuredThisWeek(c *gin.Context) {
	organization_id := c.Param("id")

	now := time.Now()
	startOfWeek := now.AddDate(0, 0, -7).Add(time.Nanosecond)

	var medicalRecords []models.MedicalRecord
	if err := models.DB.Where(("organization_id = ? AND injury_date >= ? AND injury_date < ?"), organization_id, startOfWeek, now).Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("return records from %s to %s", startOfWeek, now)
	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

func GetInjuredThisMonth(c *gin.Context) {
	organizationId := c.Param("id")
	var medicalRecords []models.MedicalRecord

	startOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	if err := models.DB.Where("organization_id = ? AND injury_date >= ? AND injury_date < ?", organizationId, startOfMonth, endOfMonth).Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

func GetRecordByStatus(c *gin.Context){
	organizationId := c.Param("id")
	status := c.Param("status")

	statusMap := map[string]string{
		"rehabilitation": "リハビリ",
		"return": "復帰済",
		"partial": "部分復帰",
		"confirm": "確認中",
		"other": "その他",
	}

	japaneseStatus, ok := statusMap[status]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	var medicalRecords []models.MedicalRecord
	if err := models.DB.Where("organization_id = ? AND status = ?", organizationId, japaneseStatus).Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

func GetRecoverThisWeek(c *gin.Context) {
	organizationId := c.Param("id")
	var medicalRecords []models.MedicalRecord

	now := time.Now()
	// 今週の開始日を計算（週の開始を月曜日とする）
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6 // 日曜日の場合、前週の月曜日に戻るために
	}
	startOfWeek := time.Date(now.Year(), now.Month(), now.Day()+offset, 0, 0, 0, 0, now.Location())

	// 今週の終了日を計算（日曜日の23:59:59）
	endOfWeek := startOfWeek.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)


	if err := models.DB.Where("organization_id = ? AND recovery_date >= ? AND recovery_date < ?", organizationId, startOfWeek, endOfWeek).Order("recovery_date ASC").Preload("Player").Find(&medicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}

func GetFilteredByCategory(c *gin.Context) {
	orgniazationId := c.Param("id")
	categoryId := c.Param("category_id")
	optionId := c.Param("option_id")

	var players []models.Player
	var playerCategories []models.PlayerCategory

	if err := models.DB.Where("category_id = ? AND category_option_id = ?", categoryId, optionId).Find(&playerCategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, pc := range playerCategories {
		if err := models.DB.Where("ID = ? AND organization_id = ?", pc.PlayerID, orgniazationId).Find(&players).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var medicalRecords []models.MedicalRecord
	for _, player := range players {
		if err := models.DB.Where("player_id = ?", player.ID).Preload("Player").Find(&medicalRecords).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"medicalRecords": medicalRecords})
}