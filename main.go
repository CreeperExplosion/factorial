package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strconv"
)

type computationRange struct {
	from int
	to   int
}

var counter int

func main() {

	strNums := os.Args[1:]
	var nums []int

	fmt.Printf("running with %d threads \n \n", runtime.NumCPU())

	for _, str := range strNums {

		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println("(" + str + ") is not a string")
			os.Exit(-1)
		}

		nums = append(nums, num)
	}

	for _, num := range nums {
		fmt.Println(fmt.Sprintf("\n\n%d! = %s \n", num, doFactorial(int64(num))))
	}
}

func doFactorial(num int64) string {
	var jobs []chan big.Int

	trds := int64(runtime.NumCPU())

	if num < 100 {
		trds = 1
	}

	facrange := num / trds

	for i := int64(0); i < trds; i++ {
		job := make(chan big.Int, 1)
		jobs = append(jobs, job)
		counter++
		go doPartialFac(job, (facrange*i)+1, facrange*(i+1))
	}

	result := big.NewInt(1)
	for _, job := range jobs {

		res := <-job
		result.Mul(result, &res)
		counter--
		fmt.Printf("%.1f %% \n", (1-float64(counter)/float64(trds))*100)
		close(job)
	}

	return result.String()
}

func doPartialFac(job chan<- big.Int, from, to int64) {
	product := big.NewInt(1)
	product.MulRange(from, to)
	job <- *product

}
