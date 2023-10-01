package main

import (
	"fmt"
	"sort"
)

func top(total map[string][]Payment) map[string][]Payment {
	topPerCat := make(map[string][]Payment, 0)
	for k, v := range total {

		fmt.Printf("The cat %v has a length of %v\n", k, len(v))
		sort.Slice(v, func(i, j int) bool {
			return v[i].Amount.Amount() > v[j].Amount.Amount()
		})

		fmt.Printf("The greatest in category %v is %v with %v\n", k, v[0], v[0].Amount.Display())

		add(len(v), 0, k, topPerCat, v)
	}
	return topPerCat
}

func add(len int, check int, k string, m map[string][]Payment, p []Payment) {
	if check < 3 {
		if len > check {
			m[k] = append(m[k], p[check])
			add(len, check+1, k, m, p)
		}
	}
}
