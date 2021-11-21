package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"project-kp/config"
	"project-kp/model"
)

func GetAllUom(c *gin.Context) {
	var list []model.Uom
	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}

func CreateUom(c *gin.Context) {
	var obj model.Uom
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&obj); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, obj)
}

func DeleteUom(c *gin.Context) {
	id := c.Param("id")
	var obj model.Uom

	// Record Not Found
	result := config.DB.First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.Uom{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success delete", "data": obj})
	}
}
