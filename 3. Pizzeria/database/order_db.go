package database

import (
	"encoding/gob"
	"os"
	"pizzeria/models"
)

const (
	OrderHistoryFile = "order_history.gob"
)

// StoreOrdersHistory encode l'historique des commandes en GOB et les stocke dans un fichier
func StoreOrdersHistory(history models.OrdersHistory) error {
	// Ouvrir le fichier en mode écriture
	file, err := os.Create(OrderHistoryFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Créer un nouvel encodeur GOB
	encoder := gob.NewEncoder(file)

	// Encoder l'historique des commandes et l'écrire dans le fichier
	if err := encoder.Encode(history); err != nil {
		return err
	}

	return nil
}

// LoadOrdersHistory charge l'historique des commandes à partir d'un fichier GOB
func LoadOrdersHistory() (models.OrdersHistory, error) {
	var history models.OrdersHistory

	// Ouvrir le fichier en mode lecture
	file, err := os.Open(OrderHistoryFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Créer un nouveau décodeur GOB
	decoder := gob.NewDecoder(file)

	// Décoder l'historique des commandes depuis le fichier
	if err := decoder.Decode(&history); err != nil {
		return nil, err
	}

	return history, nil
}

// GetLastOrder récupère la dernière commande passée
func GetLastOrder() (models.Order, error) {
	history, err := LoadOrdersHistory()
	if err != nil {
		return models.Order{}, err
	}
	return history[len(history)-1], nil
}

// AddOrder ajoute une commande à l'historique
func AddOrder(order models.Order) error {
	history, err := LoadOrdersHistory()
	if err != nil {
		return err
	}
	history = append(history, order)
	return StoreOrdersHistory(history)
}
