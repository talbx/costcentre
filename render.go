package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
)

func generatePieItems(s []Summarized) []opts.PieData {
	items := make([]opts.PieData, 0)
	for _, v := range s {
		items = append(items, opts.PieData{Name: v.Receiver + v.Sum.Display(), Value: v.Amount()})
	}
	fmt.Println(items)
	return items
}

func pieBase(total TotalSummarized) *charts.Pie {

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Enterable: true,
		}),
		charts.WithAnimation(),
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Total spendings: %v", total.TotalSum.Display()), TitleStyle: &opts.TextStyle{Padding: "50px"}, Top: "20%"}),
	)
	pie.AddSeries("pie", generatePieItems(total.Transactions), charts.WithPieChartOpts(opts.PieChart{
		Radius:   []string{"50%", "75%"},
		RoseType: "area",
		Center:   []string{"50%", "50%"},
	}))
	return pie
}

func create(total TotalSummarized) {
	page := components.NewPage()
	page.AddCharts(
		pieBase(total),
	)
	f, err := os.Create("pie.html")
	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		panic(err)
	}
}
