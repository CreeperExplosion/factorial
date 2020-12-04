package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

func main() {

	strNums := os.Args[1:]
	var nums []int

	for _, str := range strNums {

		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println("(" + str + ") is not a string")
			os.Exit(-1)
		}

		nums = append(nums, num)
	}

	for _, num := range nums {
		fmt.Println(doFactorial(int64(num)) + "\n")
	}
}

func doFactorial(num int64) string {

	var wg sync.WaitGroup

	a := big.NewInt(1)
	b := big.NewInt(1)

	wg.Add(1)
	go func() {
		a.MulRange(1, num/2)

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		b.MulRange((num/2)+1, num)
		wg.Done()
	}()

	wg.Wait()
	result := big.NewInt(1)
	result.Mul(a, b)

	return fmt.Sprintf("%d! = %s ", num, result.String())
}
