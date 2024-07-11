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

func PrintTable(table [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	for _, line := range table {
		fmt.Fprintln(writer, strings.Join(line, "\t")+"\t")
	}

	writer.Flush()
}
