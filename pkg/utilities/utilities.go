package utilities

import "time"

//Функция установки даты и времени транзакции в формате UTC (отброс лишних фозвращаемых параметров)
func TransactionDateTime (myDateTime string) (result time.Time) {
	result, _ = time.Parse("01/02/2006 3:04:05 PM", myDateTime)
	return result
}