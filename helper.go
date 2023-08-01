package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

func readCsvFile(filePath string) []Transaction {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	clients := make([]Transaction, 0)

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';' // This is our separator now
		return r
	})

	if err := gocsv.UnmarshalFile(f, &clients); err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return clients
}

func readYamlFile(r string) temp {
	viper.SetConfigName(r)      // name of config file (without extension)
	viper.AddConfigPath(".")    // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var conf temp
	err = viper.Unmarshal(&conf)
	return conf
}
