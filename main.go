package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"strconv"
	"strings"
)

type Transaction struct {
	Receiver    string `csv:"Auftraggeber"`
	Description string `csv:"Buchungstext"`
	Amount      string `csv:"Betrag"`
	Valuta      string `csv:"Valuta"`
	Buchung     string `csv:"Buchung"`
	Purpose     string `csv:"Verwendungszweck"`
	Saldo       string `csv:"Saldo"`
	Curr2       string `csv:"Curr2"`
	Curr1       string `csv:"Curr1"`
}

type Payment struct {
	Amount *money.Money
	Date   string
}

func main() {
	records := readCsvFile("data.csv")
	list := readYamlFile("r")
	payments := make(map[string][]Payment, 0)
	for _, record := range records {
		if strings.Contains(record.Amount, "-") {
			found := false
			realDeal := strings.Split(record.Amount, "-")[1]
			r1 := strings.ReplaceAll(realDeal, ".", "")
			r2 := strings.ReplaceAll(r1, ",", ".")
			f, err := strconv.ParseFloat(r2, 32)
			if err != nil {
				panic(err)
			}
			for _, receiver := range list {
				//fmt.Printf("Checking if %v is in %v\n", receiver, record.Receiver)
				if strings.Contains(strings.ToLower(record.Receiver), strings.ToLower(receiver)) || (strings.Contains(strings.ToLower(record.Purpose), strings.ToLower(receiver))) {
					payments[receiver] = append(payments[receiver], Payment{money.NewFromFloat(f, money.EUR), record.Buchung})
					found = true
					break
				}
			}
			if found == false {
				fmt.Printf("nothing found for %v, so will add %v to rest\n", record.Receiver, record.Amount)
				payments["rest"] = append(payments["rest"], Payment{
					Amount: money.NewFromFloat(f, money.EUR),
					Date:   record.Buchung,
				})
			}
		}
	}

	result := summarize(payments)
	total := totalize(result)
	create(total)
}

type temp struct {
	Data []string
}
