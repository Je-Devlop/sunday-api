package sunday

import "net/http"

type Handler struct {
	store stores
}

func (s *Handler) CreateScoops(c FrameworkContext) {
	var scoop Scoop
	if err := c.Bind(&scoop); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := s.store.CreateICreamScoop(scoop)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{})
}

func (s *Handler) GetSundayScoops(c FrameworkContext) {
	scoops, err := s.store.GetAllIceCreamScoops()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, scoops)
}

func (s *Handler) CreateTopping(c FrameworkContext) {
	var topping Topping
	if err := c.Bind(&topping); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err := s.store.CreateICreamTopping(topping)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{})
}

func (s *Handler) GetSundayTopping(c FrameworkContext) {
	toppings, err := s.store.GetAllIceCreamToppings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, toppings)
}
