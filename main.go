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
		fmt.Println(fmt.Sprintf("%d! = %s \n", num, doFactorial(int64(num))))
	}
}

func doFactorial(num int64) string {

	var wg sync.WaitGroup

	a := big.NewInt(1)
	b := big.NewInt(1)
	c := big.NewInt(1)
	d := big.NewInt(1)

	facRange := int64(num / 4)

	wg.Add(1)
	go func() {
		a.MulRange(1, facRange)

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		b.MulRange(facRange+1, facRange*2)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		c.MulRange((facRange*2)+1, facRange*3)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		d.MulRange((facRange*3)+1, num)
		wg.Done()
	}()

	wg.Wait()
	result := big.NewInt(1)
	result.Mul(a, b)
	result.Mul(result, c)
	result.Mul(result, d)

	return result.String()
}
