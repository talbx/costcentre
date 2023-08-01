package main

import (
	"github.com/Rhymond/go-money"
	"strconv"
	"strings"
)

func find(record Transaction, list temp, payments map[string][]Payment) {
	if strings.Contains(record.Amount, "-") {
		found := false
		realDeal := strings.Split(record.Amount, "-")[1]
		r1 := strings.ReplaceAll(realDeal, ".", "")
		r2 := strings.ReplaceAll(r1, ",", ".")
		f, err := strconv.ParseFloat(r2, 32)
		if err != nil {
			panic(err)
		}

		for k, items := range list.Data {
			for _, item := range items {
				if isInList(record, item) {
					payments[k] = append(payments[k], Payment{money.NewFromFloat(f, money.EUR), record.Buchung, record.Receiver})
					found = true
					break
				}
			}
		}
		if found == false {
			payments["rest"] = append(payments["rest"], Payment{
				Amount:   money.NewFromFloat(f, money.EUR),
				Date:     record.Buchung,
				Receiver: record.Receiver,
			})
		}
	}
}

func isInList(record Transaction, receiver string) bool {
	return strings.Contains(strings.ToLower(record.Receiver), strings.ToLower(receiver)) || (strings.Contains(strings.ToLower(record.Purpose), strings.ToLower(receiver)))
}
