package util

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintTable(data [][]interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "Row\t")
	for j := 0; j < len(data[0]); j++ {
		columnLetter := string(rune('A' + j))
		fmt.Fprintf(w, "%s\t", columnLetter)
	}
	fmt.Fprintln(w)

	for i, row := range data {
		fmt.Fprintf(w, "%d\t", i+1)

		for _, cell := range row {
			fmt.Fprintf(w, "%v\t", cell)
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
