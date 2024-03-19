package database

import (
	"encoding/gob"
	"os"
	"pizzeria/models"
)

const IngredientStockFile = "ingredient_stock.gob"

// StoreIngredientStock encode les stocks d'ingrédients en GOB et les stocke dans un fichier
func StoreIngredientStock(stock models.IngredientsStock) error {
	// Ouvrir le fichier en mode écriture
	file, err := os.Create(IngredientStockFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Créer un nouvel encodeur GOB
	encoder := gob.NewEncoder(file)

	// Encoder les stocks d'ingrédients et les écrire dans le fichier
	if err := encoder.Encode(stock); err != nil {
		return err
	}

	return nil
}

// LoadIngredientStock charge les stocks d'ingrédients à partir d'un fichier GOB
func LoadIngredientStock() (models.IngredientsStock, error) {
	var stock models.IngredientsStock

	// Ouvrir le fichier en mode lecture
	file, err := os.Open(IngredientStockFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Créer un nouveau décodeur GOB
	decoder := gob.NewDecoder(file)

	// Décoder les stocks d'ingrédients depuis le fichier
	if err := decoder.Decode(&stock); err != nil {
		return nil, err
	}

	return stock, nil
}

// GetIngredientStockByID récupère la quantité d'un ingrédient en stock
func GetIngredientStockByID(id int) (int, error) {
	stock, err := LoadIngredientStock()
	if err != nil {
		return 0, err
	}
	return stock[id], nil
}

// UpdateIngredientStock met à jour la quantité d'un ingrédient en stock
func UpdateIngredientStock(id, quantity int) error {
	stock, err := LoadIngredientStock()
	if err != nil {
		return err
	}
	stock[id] += quantity
	return StoreIngredientStock(stock)
}

