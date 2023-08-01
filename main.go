package main

import (
	"github.com/Rhymond/go-money"
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
	Amount   *money.Money
	Date     string
	Receiver string
}

func main() {
	records := readCsvFile("data.csv")
	list := readYamlFile("r")
	payments := make(map[string][]Payment, 0)
	for _, record := range records {
		find(record, list, payments)
	}

	result := summarize(payments)
	total := totalize(result)
	create(total)
}

type InputData struct {
	Groceries    []string `yaml:"groceries"`
	Services     []string `yaml:"services"`
	Insurances   []string `yaml:"insurances"`
	FoodDelivery []string `yaml:"foodDelivery"`
	Rest         []string `yaml:"misc"`
	Family       []string `yaml:"family"`
	Friends      []string `yaml:"friends"`
	Appartment   []string `yaml:"appartment"`
}

type ExportData struct {
	Groceries    []Payment `yaml:"groceries"`
	Services     []Payment `yaml:"services"`
	Insurances   []Payment `yaml:"insurances"`
	FoodDelivery []Payment `yaml:"foodDelivery"`
	Rest         []Payment `yaml:"misc"`
	Family       []Payment `yaml:"family"`
	Friends      []Payment `yaml:"friends"`
	Appartment   []Payment `yaml:"appartment"`
}

type temp struct {
	Data map[string][]string
}
