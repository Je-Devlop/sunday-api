package sunday

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Scoop struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	ImagePath string `json:"image_path" binding:"required"`
}

type SundayHandler struct {
	db *gorm.DB
}

func NewSundayHandler(db *gorm.DB) *SundayHandler {
	return &SundayHandler{db: db}
}

func (s *SundayHandler) CreateScoops(c *gin.Context) {
	var scoop Scoop
	if err := c.ShouldBind(&scoop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := s.db.Create(&scoop)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})

}
