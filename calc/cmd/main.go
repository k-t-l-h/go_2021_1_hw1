package main

import (
	"fmt"
	"github.com/ktlh/go_2021_1_hw_1/m/internal"
	"os"
)

func main() {

	result, err := internal.Calc(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("result:  ", result)

}
