package endpoint

import (
	"errors"
	"net/http"
	"project-kp/config"
	"project-kp/model"
	"project-kp/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUom(c *gin.Context) {
	var list []model.Uom
	config.DB.Find(&list)
	if err := config.DB.Find(&list).Error; err != nil {
		util.BadResponse(c, err)
		return
	}

	util.GoodResponse(c, list)
}

func CreateUom(c *gin.Context) {
	var obj model.Uom
	if err := c.ShouldBindJSON(&obj); err != nil {
		util.BadResponse(c, err)
		return
	}

	if result := config.DB.Create(&obj); result.Error != nil {
		util.BadResponse(c, result.Error)
		return
	}
	util.GoodResponse(c, obj)
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
