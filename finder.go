package main

import (
	"github.com/Rhymond/go-money"
	"log"
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
					payments[k] = appendPayment(k, record, payments, f)
					found = true
					break
				}
			}
		}
		if found == false {
			if strings.Contains(record.Purpose, "Apple Pay") {
				payments["APPLE_PAY"] = appendPayment("APPLE_PAY", record, payments, f)
				return
			}
			log.Printf("Could not match %+v with any category. will push to REST category\n", record)
			payments["rest"] = appendPayment("rest", record, payments, f)
		}
	}
}

func appendPayment(key string, record Transaction, payments map[string][]Payment, f float64) []Payment {
	return append(payments[key], Payment{
		Amount:   money.NewFromFloat(f, money.EUR),
		Date:     record.Buchung,
		Receiver: record.Receiver,
		Purpose:  record.Purpose,
	})
}

func isInList(record Transaction, receiver string) bool {
	return strings.Contains(strings.ToLower(record.Receiver), strings.ToLower(receiver))
}
