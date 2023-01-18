package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/Hilson-Alex/number-in-full/src/numberParser"
)

const colorReset = "\033[0m"

func main() {
	log.SetPrefix("\033[91mERROR: ")
	if len(os.Args) <= 1 {
		log.Fatal("Number argument missing. Please execute number-in-full.exe <number>", colorReset)
	}
	var arg, err = strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("Invalid Number. Please type a number between 0 and ", uint64(math.MaxUint64), colorReset)
	}
	fmt.Println(numberParser.GetNumberInFull(arg))
}
