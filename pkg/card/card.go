package card

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/utilities"
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


//Функция выборки транзакций внутри месяца
func SelectMonthTransactions(transactions []Transaction, startDate, endDate string) (monthSlice []Transaction) {
	for _, myTransaction := range transactions {
		if myTransaction.DateTime.After(utilities.TransactionDateTime(startDate)) && myTransaction.DateTime.Before(utilities.TransactionDateTime(endDate)) {
			monthSlice = append(monthSlice, myTransaction)
		}
	}
	return monthSlice
}

//Функция подсчета общей суммы операций в массиве транзакций
func TotalSumCalculation (transactions []Transaction) (totalSum int64) {
	totalSum = 0
	for _, myTransaction := range transactions {
		totalSum += myTransaction.Amount
	}
	return totalSum
}

//Функция построчной печати элементов слайса транзакций
func TransactionSlicePrinting(transactions []Transaction) {
	for _, myTransaction := range transactions {
		fmt.Printf("%+v\n", myTransaction)
	}
}