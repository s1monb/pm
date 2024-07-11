/*
Copyright © 2024 SIMON BJØRNØY <SIMON.BJORNOY@GMAIL.COM>
*/
package helpers

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// PrintTable prints a well-formatted table to the console
// it takes a 2D slice of strings as input
func PrintTable(table [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 32, 4, 0, ' ', 0)
	for _, line := range table {
		fmt.Fprintln(writer, strings.Join(line, "\t")+"\t")
	}

	writer.Flush()
}
