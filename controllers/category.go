package controllers

import (
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
	Options []string `json:"options"`
}

func CreateCategory(c *gin.Context) {
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organizationIdStr := c.Param("id")
	organizationId, err := strconv.ParseUint(organizationIdStr, 10, 64)
	fmt.Println(organizationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization ID"})
		return
	}

	var category models.Category
	err = category.Create(input.Name, uint(organizationId), input.Options)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Preload("Options").First(&category, category.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"category": category})
}

type CategoryResponse struct {
	ID uint `json:"id"`
	OrganizationID uint `json:"organization_id"`
	Name string `json:"name"`
	Options []CategoryOptionResponse `json:"options"`
}

type CategoryOptionResponse struct {
	ID uint `json:"id"`
	CategoryID uint `json:"category_id"`
	Value string `json:"value"`
}

func GetCategories(c *gin.Context) {
	organizationId := c.Param("id")
	var categories []models.Category

	if err := models.DB.Where("organization_id = ?", organizationId).Preload("Options").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoriesResponse []CategoryResponse
	for _, category := range categories {
		categoryResponse := CategoryResponse{
			ID:             category.ID,
			OrganizationID: category.OrganizationID,
			Name:           category.Name,
			Options:        []CategoryOptionResponse{},
		}

		for _, option := range category.Options {
			categoryResponse.Options = append(categoryResponse.Options, CategoryOptionResponse{
				ID:         option.ID,
				CategoryID: option.CategoryID,
				Value:      option.Value,
			})
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	c.JSON(http.StatusOK, gin.H{"categories": categoriesResponse})
}