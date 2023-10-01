package main

import (
	"github.com/Rhymond/go-money"
	"github.com/go-echarts/go-echarts/v2/charts"
	"log"
	"os"
)

type Transaction struct {
	Receiver    string `csv:"Auftraggeber/Empf�nger"`
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

	ConfigureLogger()
	dumps := readDumps()
	pies := make([]*charts.Pie, 0)
	for key, records := range dumps {

		list := readYamlFile("r")
		payments := make(map[string][]Payment, 0)
		for _, record := range records {
			find(record, list, payments)
		}

		result := summarize(payments)
		total := totalize(result)
		pie := create(key, total)
		pies = append(pies, pie)
	}

	createPage(pies)
}

func readDumps() map[string][]Transaction {
	dirs, err := os.ReadDir("./dumps")
	if err != nil {
		log.Fatal(err)
	}
	csvs := make(map[string][]Transaction, 0)
	for _, dir := range dirs {
		csv := readCsvFile("./dumps/" + dir.Name())
		csvs[dir.Name()] = csv
	}
	return csvs
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
