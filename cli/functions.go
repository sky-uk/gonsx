package main

import (
	"github.com/crackcomm/go-clitable"
)

// PrettyPrintMany - Pretty prints maps as tables
func PrettyPrintMany(headers []string, rows []map[string]interface{}) {
	table := clitable.New(headers)
	for _, row := range rows {
		table.AddRow(row)
	}
	table.Print()
}

// PrettyPrintSingle - Pretty prints map as tables
func PrettyPrintSingle(row map[string]interface{}) {
	clitable.PrintHorizontal(row)
}
