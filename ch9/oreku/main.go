package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "8000"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_CENTER,
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor},
	)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooter([]string{
		"", "", "427.0",
	})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
