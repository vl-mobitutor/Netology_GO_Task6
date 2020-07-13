package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/card"
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/utilities"
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
	myTransactions:= []card.Transaction{	//начало массива транзакций
		{
			Id: 1001,
			Amount:  -1003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("01/05/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5411",
		},

		{
			Id: 1002,
			Amount: -1001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("02/05/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5411",
		},

		{
			Id: 1003,
			Amount: -1002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("03/05/2020 5:09:01 PM"),
			Description: "Gastrobar",
			Status:  "Обработана",
			MccCode: "5812",
		},

		{
			Id: 1004,
			Amount:  -1005_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("04/05/2020 5:09:01 PM"),
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
			DateTime: utilities.TransactionDateTime("06/06/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1007,
			Amount:  -2004_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("07/06/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1008,
			Amount:  -2003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("08/06/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1009,
			Amount:  -2002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("09/06/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1010,
			Amount:  -2001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("10/06/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1011,
			Amount:  -3002_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("01/07/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1012,
			Amount:  -3001_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("02/07/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1013,
			Amount:  -3003_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("03/07/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1014,
			Amount:  -3004_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("04/07/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

		{
			Id: 1015,
			Amount:  -3005_00,
			Currency: "RUR",
			DateTime: utilities.TransactionDateTime("05/07/2020 5:09:01 PM"),
			Description: "Gipermarket",
			Status:  "Операция в обработке",
			MccCode: "5400",
		},

	}
	myCard.Transactions = myTransactions


	fmt.Println("ЗАДАНИЕ №1 - СОРТИРОВКА")
	fmt.Println("----------------- Несортированный массив транзакций--------------------------------------------")
	fmt.Println(myCard.Transactions)

	fmt.Println("------- Массив транзакций, отсортированный по убыванию суммы транзакции------------------------")
	card.SortByAmountDecrease(myCard.Transactions)
	fmt.Println(myCard.Transactions)

	fmt.Println("------- Массив транзакций, отсортированный по возрастанию суммы транзакции------------------------")
	card.SortByAmountIncrease(myCard.Transactions)
	fmt.Println(myCard.Transactions)

	fmt.Println("ЗАДАНИЕ №2 - ГОРУТИНЫ")

}
