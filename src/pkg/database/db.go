package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	smcfg "r3kk3/src/pkg/config"
	utils "r3kk3/src/pkg/utils"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Инициализация БД
func InitDB() *sql.DB {
	var config = smcfg.Init_config()
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DB,
	)

	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")
	time.Sleep(5 * time.Second)
	CreateWalletTable(*conn)
	CreateTransactionTable(*conn)
	InitWallets(*conn)
	DB = conn

	return conn
}

// Инициализация таблицы кошельков
func CreateWalletTable(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Wallet(
		id SERIAL PRIMARY KEY,
        Address VARCHAR(255) NOT NULL UNIQUE,
        Balance DECIMAL NOT NULL CHECK (Balance >= 0)
        );`)
	if err != nil {
		return
	}
	fmt.Println("wallet table created")
}

// Инициализация таблицы транзакций
func CreateTransactionTable(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Transaction(
		id SERIAL PRIMARY KEY,
		Date TIMESTAMP NOT NULL,
		Amount DECIMAL NOT NULL,
        Sender VARCHAR(255) NOT NULL,
        Recipient VARCHAR(255) NOT NULL
        );`)
	if err != nil {
		return
	}
	fmt.Println("transaction table created")
}

// Создание 10 случайных кошельков
func InitWallets(DB sql.DB) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	rows, err := DB.QueryContext(
		ctx, "SELECT * FROM Wallet",
	)
	if err != nil {
		log.Printf("Error %s when init wallets\n", err)
		return
	}
	if !(rows.Next()) {
		for i := 0; i < 10; i++ {
			_, err := DB.ExecContext(ctx,
				`INSERT INTO Wallet (
					Address, 
					Balance
				) 
				VALUES ($1, $2)`,
				utils.СreateAddress(),
				100,
			)
			if err != nil {
				log.Printf("Error %s when insert wallet\n", err)
				return
			}
		}
		fmt.Println("10 wallets created")
	}

}
