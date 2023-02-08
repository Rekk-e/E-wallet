package main

import (
	"database/sql"
	db "r3kk3/src/pkg/database"
	handlers "r3kk3/src/pkg/handlers"
)

var database sql.DB

func main() {

	// Подключаемся к БД
	db.InitDB()

	// Инициализируем хендлеры
	handlers.InitHandlers()

}
