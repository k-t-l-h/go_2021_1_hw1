package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/k-t-l-h/go_2021_1_hw1/uniq/internal"
	"io"
	"os"
)

var p internal.Params

func flagInit() {
	flag.BoolVar(&(p.Count), "c", false, "count the number of occurrences")
	flag.BoolVar(&(p.Double), "d", false, "print only repeated lines")
	flag.BoolVar(&(p.Uniq), "u", false, "print only unique lines")
	flag.BoolVar(&(p.Case), "i", false, "case sensitivity")

	flag.IntVar(&(p.FieldsOffset), "f", 0, "skip n fields")
	flag.IntVar(&(p.CharOffset), "s", 0, "skip n chars")
}

func main() {
	var (
		inputFlag  string
		outputFlag string
		err        error
	)
	flagInit()
	flag.Parse()

	input := os.Stdin
	output := os.Stdout

	inputFlag = flag.Arg(0)
	outputFlag = flag.Arg(1)

	if inputFlag != "" {
		input, err = os.Open(inputFlag)
		if err != nil {
			fmt.Print("Не удалось открыть файл на чтение")
			os.Exit(-1)
		}
	}

	if outputFlag != "" {
		output, err = os.Create(outputFlag)
		if err != nil {
			fmt.Print("Не удалось открыть файл на запись")
			os.Exit(-1)
		}
	}

	if !checkFlags(p.Double, p.Uniq, p.Count) {
		flag.Usage()
	}

	values := getLines(input)
	values = internal.Uniq(values, p)
	err = giveLines(output, values)
	if err != nil {
		fmt.Print("Не удалось записать в файл")
	}

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
	return !((a && b) || (a && c) || (b && c))
}
