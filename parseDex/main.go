package main

import (
	"GolangStudy/parseDex/service"
	"fmt"
	"os"
)

func main() {

	fmt.Println("start parse dex file")

	if len(os.Args) <= 1 {
		fmt.Println("please input file with command line")
		return
	}
	inputFile := os.Args[1]
	fmt.Println("input file path is: ", inputFile)
	service.ParseDex(inputFile)

}
