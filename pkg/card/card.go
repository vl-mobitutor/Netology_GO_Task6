package card

import (
	"sort"
	"time"
)

type Transaction struct {
	Id int64 // id-транзакции
	Amount int64 //сумма операции в копейках/центах
	Currency string //валюта операции
	DateTime time.Time //дата и время операции в формате UTC
	Description string //описание операции
	Status string //текущий статус операции
	MccCode string //код точки продаж по каталогу Merchant Category Code
}

type Card struct {
	Id int64 // id-карты
	Number string //номер карты
	PaymentSystem string //платежная система
	BankIssuer string //банк-эмитент карты
	Balance int64 //текущий баланс карты в копейках/центах
	Currency string //валюта карточного счета
	Transactions []Transaction
}

//Функция сортировки массива транзакций по убыванию суммы
func SortByAmountDecrease(transactions []Transaction) []Transaction{
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Amount < transactions[j].Amount
	})
	return transactions
}

//Функция сортировки массива транзакций по возрастанию суммы
func SortByAmountIncrease(transactions []Transaction) []Transaction{
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Amount > transactions[j].Amount
	})
	return transactions
}