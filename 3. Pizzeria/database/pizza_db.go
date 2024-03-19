package database

import (
	"pizzeria/models"
)

var pizzas = []models.Pizza{
	{
		ID:    1,
		Name:  "3 fromages",
		Price: 12.5,
		Ingredients: []models.Ingredient{
			models.Mozzarella,
			models.Parmesan,
			models.Gorgonzola,
		},
	},
	{
		ID:    2,
		Name:  "Reine",
		Price: 14.0,
		Ingredients: []models.Ingredient{
			models.Mozzarella,
			models.Tomate,
			models.Jambon,
			models.Champignons,
		},
	},
	{
		ID:    3,
		Name:  "4 saisons",
		Price: 13.0,
		Ingredients: []models.Ingredient{
			models.Mozzarella,
			models.Tomate,
			models.Jambon,
			models.Champignons,
			models.Artichauts,
		},
	},
}

// GetPizzas récupère la liste des pizzas disponibles depuis la base de données
func GetPizzas() ([]models.Pizza, error) {
	// Code pour récupérer les pizzas depuis la base de données
	return pizzas, nil
}

// GetPizzaByID récupère une pizza depuis la base de données
func GetPizzaByID(id int) (models.Pizza, error) {
	// Code pour récupérer une pizza depuis la base de données
	for _, pizza := range pizzas {
		if pizza.ID == id {
			return pizza, nil
		}
	}
	return models.Pizza{}, nil
}
