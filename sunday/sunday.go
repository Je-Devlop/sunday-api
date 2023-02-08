package sunday

import (
	"net/http"

	"gorm.io/gorm"
)

type Scoop struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	ImagePath string `json:"image_path" binding:"required"`
}

type stores interface {
	New(*Scoop) error
}

type SundayHandler struct {
	store stores
}

func NewSundayHandler(store stores) *SundayHandler {
	return &SundayHandler{store: store}
}

func (s *SundayHandler) CreateScoops(c FrameworkContext) {
	var scoop Scoop
	if err := c.Bind(&scoop); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := s.store.New(&scoop)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{})

}

func (s *SundayHandler) GetSundayScoops(c FrameworkContext) {

}
