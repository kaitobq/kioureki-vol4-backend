package controllers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePlayerInput struct {
	Name string `json:"name" binding:"required"`
}

func CreatePlayer(c *gin.Context) {
	var input CreatePlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organizationIdStr := c.Param("id")
	organizationId, err := strconv.ParseUint(organizationIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player := models.Player{Name: input.Name, OrganizationID: uint(organizationId)}
	if err := models.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"player": player})
}

func GetPlayers(c *gin.Context) {
	organizationId := c.Param("id")

	var players []models.Player
	if err := models.DB.Where("organization_id = ?", organizationId).Find(&players).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"players": players})
}

func DeletePlayer(c *gin.Context) {
	playerId := c.Param("player_id")

	if err := models.DB.Where("id = ?", playerId).Delete(&models.Player{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

type AddCategoryInput struct {
	CategoryID uint `json:"category_id" binding:"required"`
	CategoryOptionID uint `json:"category_option_id" binding:"required"`
}

func AddCategory(c *gin.Context) {
	playerIdStr := c.Param("player_id")
	playerId, err := strconv.ParseUint(playerIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	
	var input AddCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playerCategory models.PlayerCategory
	playerCategory.PlayerID = uint(playerId)
	playerCategory.CategoryID = input.CategoryID
	playerCategory.CategoryOptionID = input.CategoryOptionID

	if err := models.DB.Create(&playerCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playerCategory})
}

type PlayerCategoryResponse struct {
	PlayerID uint `json:"player_id"`
	PlayerName string `json:"player_name"`
	Categories []CategoryDetail `json:"categories"`
}

type CategoryDetail struct {
	CategoryID         uint   `json:"category_id"`
	CategoryName       string `json:"category_name"`
	CategoryOptionID   uint   `json:"category_option_id"`
	CategoryOptionValue string `json:"category_option_value"`
}

func GetPlayerCategories(c *gin.Context) {
	organizationId := c.Param("id")
	var players []models.Player

	if err := models.DB.Where("organization_id = ?", organizationId).Find(&players).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playerCategories []models.PlayerCategory
	for _, player := range players {
		var p_categories []models.PlayerCategory
		if err := models.DB.Where("player_id = ?", player.ID).Find(&p_categories).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		playerCategories = append(playerCategories, p_categories...)
	}

	var playerCategoriesResponse []PlayerCategoryResponse

	for _, player := range players {
		var categories []CategoryDetail
		for _, playerCategory := range playerCategories {
			if playerCategory.PlayerID == player.ID {
				var category models.Category
				if err := models.DB.Where("id = ?", playerCategory.CategoryID).First(&category).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				var option models.CategoryOption
				if err := models.DB.Where("id = ?", playerCategory.CategoryOptionID).First(&option).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				categories = append(categories, CategoryDetail{
					CategoryID:         category.ID,
					CategoryName:       category.Name,
					CategoryOptionID:   option.ID,
					CategoryOptionValue: option.Value,
				})
			}
		}
		playerCategoriesResponse = append(playerCategoriesResponse, PlayerCategoryResponse{
			PlayerID:   player.ID,
			PlayerName: player.Name,
			Categories: categories,
		})
	}

	c.JSON(http.StatusOK, gin.H{"playerCategories": playerCategoriesResponse})
}

type PlayerCategoryDetail struct {
	PlayerID uint `json:"player_id"`
	PlayerName string `json:"player_name"`
	Categories []CategoryDetail `json:"categories"`
	MedicalRecords []models.MedicalRecord `json:"medical_records"`
}

func GetPlayerDetail(c *gin.Context) {
	player_id := c.Param("player_id")
	var player models.Player
	if err := models.DB.Where(("id = ?"), player_id).First(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playerCategories []models.PlayerCategory
	if err := models.DB.Where("player_id = ?", player_id).Find(&playerCategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var categories []CategoryDetail
	for _, playerCategory := range playerCategories {
		if playerCategory.PlayerID == player.ID {
			var category models.Category
			if err := models.DB.Where("id = ?", playerCategory.CategoryID).First(&category).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var option models.CategoryOption
			if err := models.DB.Where("id = ?", playerCategory.CategoryOptionID).First(&option).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			categories = append(categories, CategoryDetail{
				CategoryID:         category.ID,
				CategoryName:       category.Name,
				CategoryOptionID:   option.ID,
				CategoryOptionValue: option.Value,
			})
		}
	}

	var medicalrecords []models.MedicalRecord
	if err := models.DB.Where("player_id = ?", player_id).Find(&medicalrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playerDetail PlayerCategoryDetail
	playerDetail.PlayerID = player.ID
	playerDetail.PlayerName = player.Name
	playerDetail.Categories = categories
	playerDetail.MedicalRecords = medicalrecords

	c.JSON(http.StatusOK, gin.H{"playerDetail": playerDetail})
}