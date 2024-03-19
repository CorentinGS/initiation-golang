package handlers

import "net/http"

func GetIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ingredients"))
}

func GetPurchaseIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Purchase ingredients"))
}

func PostPurchaseIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Purchase ingredients"))
}
