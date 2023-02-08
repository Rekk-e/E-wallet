package db

import (
	"context"
	"r3kk3/src/pkg/models"
	"time"

	utils "r3kk3/src/pkg/utils"
)

// Получение всех кошельков
func GetWallets() ([]models.Wallet, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Wallet
	var wallets []models.Wallet = nil
	rows, err := DB.QueryContext(
		ctx,
		"SELECT * FROM Wallet",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&query.Id,
			&query.Address,
			&query.Balance,
		)
		wallets = append(wallets, query)
	}

	return wallets, nil
}

// Получение баланса конкретного кошелька
func GetBalance(Address string) (models.Message, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var balance string
	err := DB.QueryRowContext(
		ctx, "SELECT Balance FROM Wallet WHERE Address = $1",
		Address,
	).Scan(&balance)

	if err != nil {
		return models.Message{Response: "Error when get balance\n"}, err
	}

	return models.Message{Response: "Your balance: " + balance + " у.е."}, nil
}

// Проверка кошельков на существование в БД
func CheckWallets(Transaction models.Transaction) (models.Message, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Wallet
	var addresses []string
	rows, err := DB.QueryContext(
		ctx,
		"SELECT * FROM Wallet WHERE Address = $1 OR Address = $2",
		Transaction.From,
		Transaction.To,
	)

	if err != nil {
		return models.Message{Response: "Error when get wallets"}, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&query.Id,
			&query.Address,
			&query.Balance,
		)
		addresses = append(addresses, query.Address)
	}
	from_exist := utils.StringInSlice(Transaction.From, addresses)
	to_exist := utils.StringInSlice(Transaction.To, addresses)
	if !(from_exist) && !(to_exist) {
		return models.Message{Response: "Wallets does not exists"}, err
	} else if !(from_exist) {
		return models.Message{Response: "The sender's wallet does not exist"}, err
	} else if !(to_exist) {
		return models.Message{Response: "Recipient's wallet does not exist"}, err
	}

	return models.Message{Response: "Success"}, err
}

// Перевод денежных средств между двумя кошельками
func TransferMoney(Transaction models.Transaction) (models.Message, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(
		ctx, `UPDATE Wallet SET 
		Balance = Balance - $1
		WHERE Address = $2`,
		Transaction.Amount,
		Transaction.From,
	)

	if err != nil {
		return models.Message{Response: "Not enough money"}, err
	}

	_, err = DB.ExecContext(
		ctx, `UPDATE Wallet SET 
		Balance = Balance + $1
		WHERE Address = $2`,
		Transaction.Amount,
		Transaction.To,
	)

	if err != nil {
		return models.Message{Response: "Error when transfer money"}, err
	}
	return models.Message{Response: "Success"}, err
}

// Отправка денежных средств
func Send(Transaction models.Transaction) (models.Message, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	result, err := CheckWallets(Transaction)
	if result.Response != "Success" {
		return result, err
	}

	result, err = TransferMoney(Transaction)
	if result.Response != "Success" {
		return models.Message{Response: "Not enough money"}, err
	}

	_, err = DB.ExecContext(ctx,
		`INSERT INTO Transaction (
			Date,
			Amount, 
			Sender, 
			Recipient) 
		VALUES ($1, $2, $3, $4)`,
		time.Now(),
		Transaction.Amount,
		Transaction.From,
		Transaction.To,
	)

	if err != nil {
		return models.Message{Response: "Error when create transaction"}, err
	}
	return models.Message{Response: "Transaction successful"}, nil
}

// Получение последних транзакций
func GetLast(count int) ([]models.Transaction, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var query models.Transaction
	var transactions []models.Transaction
	rows, err := DB.QueryContext(
		ctx,
		"SELECT * FROM Transaction ORDER BY Date DESC LIMIT $1",
		count,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&query.Id,
			&query.Date,
			&query.Amount,
			&query.From,
			&query.To,
		)
		transactions = append(transactions, query)
	}

	return transactions, nil
}
