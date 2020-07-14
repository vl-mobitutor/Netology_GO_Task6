package card

import (
	"github.com/vl-mobitutor/Netology_GO_Task6/pkg/utilities"
	"reflect"
	"testing"
)


func TestSortByAmountDecrease(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{
			name: "Тест сортировки массива по убыванию суммы транзакции",
			args: args{
				[]Transaction{
					{
						Id:          1001,
						Amount:      -1003_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/01/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1002,
						Amount:      -1001_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/02/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1003,
						Amount:      -1002_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/03/2020 5:09:01 PM"),
						Description: "Gastrobar",
						Status:      "Обработана",
						MccCode:     "5812",
					},

					{
						Id:          1004,
						Amount:      -1005_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/04/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1005,
						Amount:      -1004_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/05/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5400",
					},
				},
			},
			want: []Transaction{
				{
					Id:          1004,
					Amount:      -1005_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/04/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
				{
					Id:          1005,
					Amount:      -1004_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/05/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5400",
				},
				{
					Id:          1001,
					Amount:      -1003_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/01/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
				{
					Id:          1003,
					Amount:      -1002_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/03/2020 5:09:01 PM"),
					Description: "Gastrobar",
					Status:      "Обработана",
					MccCode:     "5812",
				},
				{
					Id:          1002,
					Amount:      -1001_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/02/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
			},
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByAmountDecrease(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByAmountDecrease() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestSortByAmountIncrease(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{
			name: "Тест сортировки массива по возрастанию суммы транзакции",
			args: args{
				[]Transaction{
					{
						Id:          1001,
						Amount:      -1003_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/01/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1002,
						Amount:      -1001_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/02/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1003,
						Amount:      -1002_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/03/2020 5:09:01 PM"),
						Description: "Gastrobar",
						Status:      "Обработана",
						MccCode:     "5812",
					},

					{
						Id:          1004,
						Amount:      -1005_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/04/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5411",
					},

					{
						Id:          1005,
						Amount:      -1004_00,
						Currency:    "RUR",
						DateTime:    utilities.TransactionDateTime("05/05/2020 5:09:01 PM"),
						Description: "Gipermarket",
						Status:      "Операция в обработке",
						MccCode:     "5400",
					},
				},
			},
			want: []Transaction{
				{
					Id:          1002,
					Amount:      -1001_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/02/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
				{
					Id:          1003,
					Amount:      -1002_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/03/2020 5:09:01 PM"),
					Description: "Gastrobar",
					Status:      "Обработана",
					MccCode:     "5812",
				},
				{
					Id:          1001,
					Amount:      -1003_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/01/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
				{
					Id:          1005,
					Amount:      -1004_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/05/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5400",
				},
				{
					Id:          1004,
					Amount:      -1005_00,
					Currency:    "RUR",
					DateTime:    utilities.TransactionDateTime("05/04/2020 5:09:01 PM"),
					Description: "Gipermarket",
					Status:      "Операция в обработке",
					MccCode:     "5411",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByAmountIncrease(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByAmountIncrease() = %v, want %v", got, tt.want)
			}
		})
	}
}