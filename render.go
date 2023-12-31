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
		items = append(items, opts.PieData{Name: v.Category + v.Sum.Display(), Value: v.Amount()})
	}
	return items
}

func pieBase(name string, total TotalSummarized) *charts.Pie {

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Enterable: true,
		}),
		charts.WithAnimation(),
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Total spendings in %v: %v", name, total.TotalSum.Display()), TitleStyle: &opts.TextStyle{Padding: "50px"}, Top: "20%"}),
	)
	pie.AddSeries(name, generatePieItems(total.Transactions), charts.WithPieChartOpts(opts.PieChart{
		Radius:   []string{"50%", "75%"},
		RoseType: "area",
		Center:   []string{"50%", "50%"},
	}))
	return pie
}

func barBase(total TotalSummarized) *charts.Bar {

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Enterable: true,
		}),
		charts.WithAnimation(),
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Total spendings per category: %v", total.TotalSum.Display()), TitleStyle: &opts.TextStyle{Padding: "50px"}, Top: "20%"}),
	)
	arr := make([]string, 0)
	for _, tx := range total.Transactions {
		arr = append(arr, tx.Category)
	}

	axis := bar.SetXAxis(arr)
	bd := make([]opts.BarData, 0)
	for _, tx := range total.Transactions {
		bd = append(bd, opts.BarData{Name: tx.Category + ": " + tx.Sum.Display(), Value: tx.Sum.Absolute()})
	}
	axis.AddSeries("Categories", bd)
	return bar
}

func fineGrained(total TotalSummarized, axis *charts.Bar) {
	for _, tx := range total.Transactions {
		bd := make([]opts.BarData, 0)
		for _, payment := range tx.Payments {
			bd = append(bd, opts.BarData{Name: payment.Receiver, Value: payment.Amount.Amount()})
		}
		axis.AddSeries(tx.Category, bd)
	}
}

func create(name string, total TotalSummarized) *charts.Pie {
	return pieBase(name, total)
}

func createPage(month string, charts []*charts.Pie) *os.File {
	page := components.NewPage()
	for _, chart := range charts {
		page.AddCharts(
			chart,
		)
	}
	f, err := os.Create(fmt.Sprintf("%v.html", month))
	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		panic(err)
	}

	return f

}
