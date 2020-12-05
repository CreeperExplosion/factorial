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

	args := os.Args[1:]

	fmt.Printf("running with %d threads \n \n", runtime.NumCPU())

	num, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("%d is not a number", num)
		return
	}

	fi, err := os.OpenFile(args[1], os.O_WRONLY|os.O_CREATE|os.O_APPEND, 777)

	if err != nil {
		fmt.Println(err)
		return
	}

	out := doFactorial(int64(num))

	_, err = fi.WriteString(out)

	if err != nil {
		fmt.Println(err)
		return
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
