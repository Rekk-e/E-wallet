package models

type Transaction struct {
	Id     int
	Date   string
	Amount float32
	From   string
	To     string
}
