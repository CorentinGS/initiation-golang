package handlers

import "net/http"

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Orders"))
}

func PostOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create order"))
}
