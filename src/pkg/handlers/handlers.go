package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	commands "r3kk3/src/pkg/database"
	models "r3kk3/src/pkg/models"
	utils "r3kk3/src/pkg/utils"

	"github.com/gorilla/mux"
)

// Запрос на отправку денежных средств
func Send(w http.ResponseWriter, r *http.Request) {
	var new_transaction models.Transaction
	json.NewDecoder(r.Body).Decode(&new_transaction)
	response, _ := commands.Send(new_transaction)
	jsonResponse, jsonError := json.Marshal(response)
	utils.ReturnJson(w, jsonResponse, jsonError)

}

// Запрос на получение последних транзакций
func GetLast(w http.ResponseWriter, r *http.Request) {
	var get_last models.GetLast
	json.NewDecoder(r.Body).Decode(&get_last)
	transactions, err := commands.GetLast(get_last.Count)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Message{Response: "Get transactions failed"})
		utils.ReturnJson(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(transactions)
		utils.ReturnJson(w, jsonResponse, jsonError)
	}
}

// Запрос на получение баланса
func GetBalance(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	response, _ := commands.GetBalance(address)
	jsonResponse, jsonError := json.Marshal(response)
	utils.ReturnJson(w, jsonResponse, jsonError)

}

// Запрос на получение кошельков (в тз небыло, добавил для удобства)
func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := commands.GetWallets()
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Message{Response: "Get wallets failed"})
		utils.ReturnJson(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(wallets)
		utils.ReturnJson(w, jsonResponse, jsonError)
	}
}

// Инициализация хендлеров
func InitHandlers() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/send", Send).Methods("POST")
	router.HandleFunc("/api/transactions", GetLast).Methods("GET")
	router.HandleFunc("/api/wallet/{address}/balance", GetBalance).Methods("GET")
	router.HandleFunc("/api/wallets", GetWallets).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
