package sunday

type Scoop struct {
	Name      string `json:"name" binding:"required"`
	ImagePath string `json:"image_path" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
}

type Topping struct {
	Name      string `json:"name" binding:"required"`
	ImagePath string `json:"image_path" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
}

type stores interface {
	CreateICreamScoop(Scoop) error
	GetAllIceCreamScoops() ([]Scoop, error)
	CreateICreamTopping(Topping) error
	GetAllIceCreamToppings() ([]Topping, error)
}

func NewSundayHandler(store stores) *Handler {
	return &Handler{store: store}
}
