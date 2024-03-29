package internal

import (
	"strconv"
	"strings"
)

func Uniq(values []string, params Params) []string {
	table, count := getUnique(values, params.Case, params.FieldsOffset, params.CharOffset)
	//показываем числа или нет
	if params.Count {
		return showCount(table, count)
	}
	if params.Double {
		return showUniqless(table, count)
	} else if params.Uniq {
		return showUniq(table, count)
	} else {
		return table
	}
}

func getUnique(table []string, caseFlag bool, FieldsOffset, CharOffset int) ([]string, []int) {
	if len(table) == 0 {
		return table, nil
	}

	uniqueTable := make([]string, 0)
	uniqueCount := make([]int, 0)
	count := 1

	if len(table) > 0 {
		uniqueTable = append(uniqueTable, table[0])
	} else {
		uniqueTable = table
	}

	check := func(input string) string {

		str := strings.Split(input, " ")
		if FieldsOffset < len(str) {
			str = str[FieldsOffset:]
		}
		input = strings.Join(str, " ")
		if CharOffset != 0 && CharOffset < len(input) {
			return input[CharOffset:]
		}
		return input
	}
	if caseFlag {
		check = func(input string) string {
			str := strings.Split(input, " ")
			if FieldsOffset < len(str) {
				str = str[FieldsOffset:]
			}
			input = strings.Join(str, " ")
			if CharOffset != 0 && CharOffset < len(input) {
				return strings.ToLower(input[CharOffset:])
			}
			return strings.ToLower(input)
		}
	}

	for i := 1; i < len(table); i++ {
		if check(table[i]) != check(table[i-1]) {
			uniqueTable = append(uniqueTable, table[i])
			uniqueCount = append(uniqueCount, count)
			count = 1
		} else {
			count++
		}
	}
	uniqueCount = append(uniqueCount, count)
	return uniqueTable, uniqueCount
}

func showCount(table []string, count []int) []string {

	for i := 0; i < len(table); i++ {
		table[i] = strconv.Itoa(count[i]) + " " + table[i]
	}
	return table
}

func showUniq(table []string, count []int) []string {
	uniq := make([]string, 0, len(table))
	for i := 0; i < len(table); i++ {
		if count[i] == 1 {
			uniq = append(uniq, table[i])
		}
	}
	return uniq
}

func showUniqless(table []string, count []int) []string {
	uniqless := make([]string, 0, len(table))

	for i := 0; i < len(table); i++ {
		if count[i] > 1 {
			uniqless = append(uniqless, table[i])
		}
	}
	return uniqless
}
