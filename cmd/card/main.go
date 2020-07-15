package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/card"
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/utilities"
	"sync"

	//"sync"
)

func main() {

	//Инициализация данных карты
	myCard := &card.Card{
		Id:1,
		Number: "1111 2222 3333 4444",
		PaymentSystem: "Master",
		BankIssuer: "SuperBank",
		Balance: 100_000_00,
		Currency: "RUR",

	}

	//Инициализация массива транзакций
	myCard.Transactions = []card.Transaction{	//начало массива транзакций
		{
			Id: 1001,
			Amount:  -1003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/01/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5411",
		},

		{
			Id: 1002,
			Amount: -1001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/02/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5411",
		},

		{
			Id: 1003,
			Amount: -1002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/03/2020 5:09:01 PM"),
			Description: "Gastrobar",
			Status:  "Обработана",
			MccCode: "5812",
		},

		{
			Id: 1004,
			Amount:  -1005_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/04/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5411",
		},

		{
			Id: 1005,
			Amount:  -1004_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/05/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1006,
			Amount:  -2005_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("06/01/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1007,
			Amount:  -2004_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("06/02/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1008,
			Amount:  -2003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("06/03/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1009,
			Amount:  -2002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("06/04/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1010,
			Amount:  -2001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("06/05/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1011,
			Amount:  -3002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/01/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1012,
			Amount:  -3001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/02/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1013,
			Amount:  -3003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/03/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1014,
			Amount:  -3004_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/04/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1015,
			Amount:  -3005_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/05/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

	}


	fmt.Println("ЗАДАНИЕ №1 - СОРТИРОВКА")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("Несортированный массив транзакций:")
	fmt.Println("-------------------------------------------------------------------")
	card.TransactionSlicePrinting(myCard.Transactions)
	fmt.Println()

	fmt.Println("Массив транзакций, отсортированный по убыванию суммы транзакции:")
	fmt.Println("-------------------------------------------------------------------")
	card.SortByAmountDecrease(myCard.Transactions)
	card.TransactionSlicePrinting(myCard.Transactions)
	fmt.Println()

	fmt.Println("Массив транзакций, отсортированный по возрастанию суммы транзакции:")
	fmt.Println("-------------------------------------------------------------------")
	card.SortByAmountIncrease(myCard.Transactions)
	card.TransactionSlicePrinting(myCard.Transactions)
	fmt.Println()


	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("ЗАДАНИЕ №2 - ГОРУТИНЫ")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	fmt.Println("Подготовка данных - печать map'а слайсов транзакций, сгруппированных по месяцам:")
	selectedTransactionsMap :=card.SelectMonthTransactions(myCard.Transactions)
	card.TransactionMapPrinting(selectedTransactionsMap)

	wg := sync.WaitGroup{}
	wg.Add(len(selectedTransactionsMap))
	goroutineCounter := 0
	for month, monthlyTransactions := range selectedTransactionsMap {
		goroutineCounter++
		go func(myMonth string, myTransactions []card.Transaction, myCounter int) {
			totalSum := card.TotalSumCalculation(myTransactions)
			fmt.Printf("\nГорутина %d - Сумма транзакций за %s составила: %d копеек", myCounter, myMonth, totalSum)
			wg.Done()
		}(month, monthlyTransactions, goroutineCounter)
	}
	wg.Wait()
	fmt.Printf("\nРабота горутин завершена!")
}
