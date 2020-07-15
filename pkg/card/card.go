package card

import (
	"fmt"
	"sort"
	"strconv"
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


//Функция "нарезки" исходного массива транзакций в  map из слайсов транзакций, сгруппированных по месяцам
func SelectMonthTransactions(transactions []Transaction)  (selectedTransactions map[string][]Transaction){
	selectedTransactions = map[string][]Transaction{}
	for _, myTransaction := range transactions {
		myMonth := myTransaction.DateTime.Month().String() + "_" + strconv.Itoa(myTransaction.DateTime.Year())
		selectedTransactions[myMonth] = append(selectedTransactions[myMonth], myTransaction)
		}
	return selectedTransactions
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


//Функция построчной печати элементов map'a транзакций
func TransactionMapPrinting(transactions map[string][]Transaction) {
	for key, value := range transactions {
		fmt.Println(key, value)
	}
}