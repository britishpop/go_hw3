package main

import (
	"fmt"
	"time"

	"go_hw3/pkg/card"
)

func main() {
	t1 := &card.Transaction{
		Id:     0,
		Type:   "debit",
		Sum:    735_55,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC).Unix(),
	}
	t2 := &card.Transaction{
		Id:     1,
		Type:   "refill",
		Sum:    2_000_00,
		Status: "done",
		Date:   time.Date(2020, time.June, 11, 1, 46, 40, 0, time.UTC).Unix(),
	}
	t3 := &card.Transaction{
		Id:     2,
		Type:   "debit",
		Sum:    1203_91,
		Status: "processing",
		MCC:    "5812",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC).Unix(),
	}

	master := &card.Card{
		Id:           0,
		Issuer:       "MasterCard",
		Balance:      45_000_00,
		Currency:     "RUB",
		Number:       "0808",
		Icon:         "https://dlpng.com/png/6794578",
		Transactions: []card.Transaction{*t1, *t2, *t3},
	}
	fmt.Println("main card: ", master)
	fmt.Println("start transactions: ", master.Transactions)

	newTr := &card.Transaction{
		Id:     3,
		Type:   "debit",
		Sum:    2560_00,
		Status: "processing",
		MCC:    "5812",
		Date:   time.Date(2020, time.June, 15, 14, 30, 10, 0, time.UTC).Unix(),
	}
	card.AddTransaction(master, newTr)
	fmt.Println("transactions after add: ", master.Transactions)

	mcc := []string{"5812", "5999", "0000"}
	sum := card.SumByMCC(master.Transactions, mcc)
	fmt.Println("sum transactions with mcc", mcc, ":", sum)

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println("mcc category of code", master.Transactions[0].MCC, ":", category)

	otherCategory := card.TranslateMCC("5999")
	fmt.Println("mcc category of code 5999:", otherCategory)
}
