package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
)

type Summarized struct {
	Category string
	Payments []Payment
	Sum      *money.Money
}

func (s Summarized) Amount() int64 {
	return s.Sum.Amount()
}

type TotalSummarized struct {
	TotalSum     *money.Money
	Transactions []Summarized
}

func summarize(e map[string][]Payment) []Summarized {
	s := make([]Summarized, 0)
	for k, v := range e {
		sum := money.NewFromFloat(0.0, money.EUR)
		for _, payment := range v {
			x, err := sum.Add(payment.Amount)
			if err != nil {
				panic(err)
			}
			sum = x
		}

		s = append(s, Summarized{
			Category: k,
			Payments: v,
			Sum:      sum,
		})

		fmt.Printf("Cat:%v, Sum: %v\n", k, sum.Display())
	}
	return s
}

func totalize(s []Summarized) TotalSummarized {
	n := money.NewFromFloat(0, money.EUR)
	for _, su := range s {
		y, err := n.Add(su.Sum)
		if err != nil {
			panic(err)
		}
		n = y
	}
	return TotalSummarized{TotalSum: n, Transactions: s}
}
