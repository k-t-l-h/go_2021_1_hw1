package main

import (
	"bufio"
	"flag"
	"github.com/k-t-l-h/go_2021_1_hw1/uniq/internal"
	"io"
	"log"
	"os"
)

var inputFlag string
var outputFlag string

func flagInit(p internal.Params) {
	flag.BoolVar(&(p.Count), "c", false, "count the number of occurrences")
	flag.BoolVar(&(p.Double), "d", false, "print only repeated lines")
	flag.BoolVar(&(p.Uniq), "u", false, "print only unique lines")
	flag.BoolVar(&(p.Case), "i", false, "case sensitivity")
	flag.StringVar(&inputFlag, "in", "", "input to file")
	flag.StringVar(&outputFlag, "out", "", "output to file")
}

func main() {
	var p internal.Params
	flagInit(p)
	flag.Parse()

	log.Print(p)
	input := os.Stdin
	output := os.Stdout

	if inputFlag != "" {
		input, _ = os.Open(inputFlag)
	}

	if outputFlag != "" {
		output, _ = os.Open(outputFlag)
	}

	if !checkFlags(true, true, false) {
		flag.Usage()
	}

	values := getLines(input)
	values = internal.Uniq(values, p)
	log.Print(values)
	log.Print(giveLines(output, values))

}

func getLines(input io.Reader) []string {
	buf := bufio.NewScanner(input)

	table := make([]string, 0)
	for buf.Scan() {
		table = append(table, buf.Text())
	}
	return table
}

func giveLines(output io.Writer, table []string) error {
	for _, line := range table {
		_, err := io.WriteString(output, line+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func checkFlags(a, b, c bool) bool {
	return !(a&& b) && !(a&&c) && !(b&&c) && (!(a && b && c) || !(a || b || c))
}
