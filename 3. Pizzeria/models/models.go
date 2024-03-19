package models

// Ingredient représente un ingrédient
type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Mozzarella  = Ingredient{ID: 1, Name: "Mozzarella"}
	Parmesan    = Ingredient{ID: 2, Name: "Parmesan"}
	Gorgonzola  = Ingredient{ID: 3, Name: "Gorgonzola"}
	Tomate      = Ingredient{ID: 4, Name: "Tomate"}
	Jambon      = Ingredient{ID: 5, Name: "Jambon"}
	Champignons = Ingredient{ID: 6, Name: "Champignons"}
	Œuf         = Ingredient{ID: 7, Name: "Œuf"}
	Olives      = Ingredient{ID: 8, Name: "Olives"}
	Artichauts  = Ingredient{ID: 9, Name: "Artichauts"}
	Anchois     = Ingredient{ID: 10, Name: "Anchois"}
)

// IngredientsStock représente les stocks d'ingrédients
type IngredientsStock map[int]int

// Pizza représente une pizza
type Pizza struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Price       float64      `json:"price"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Order représente une commande de pizza
type Order struct {
	PizzaID  int    `json:"pizza_id"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

type OrdersHistory []Order

// Purchase représente un achat de stocks d'ingrédients
type Purchase struct {
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}
