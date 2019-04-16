package pg01

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type intslice []int

// Now, for our new type, implement the two methods of
// the flag.Value interface...
// The first method is String() string
func (i *intslice) String() string {
	return fmt.Sprintf("%d", *i)
}

// The second method is Set(value string) error
func (i *intslice) Set(value string) error {

	str := strings.Split(value, ",")
	for _, v := range str {
		intv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*i = append(*i, intv)
	}
	return nil
}

// Lenght : getting length of slice
func (i *intslice) Length() int {
	return len(*i)
}

func main() {
	var fSum, fMultiply intslice
	fFibo := flag.Int("fibo", 0, "generate fibonnaci series by N , Example : '--fibo=4'")
	fPrime := flag.Int("prime", 0, "generate prime series by N , Example '--prime=4' ")
	flag.Var(&fSum, "sum", "sumary of your integer list , example : '--sum=3,2...' ")
	flag.Var(&fMultiply, "multiply", "multiply of your integer list, example : '--multiply=3,2...' ")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
	} else {
		if fSum.Length() > 0 {
			fmt.Printf("the sum of %v is %v", fSum.String(), Sum(fSum))
		}
		if fMultiply.Length() > 0 {
			fmt.Printf("the multiply of %v is %v", fMultiply.String(), Multiply(fMultiply))
		}
		if *fFibo > 0 {
			fmt.Printf("the first N fibbonaci series of %v is %v", *fFibo, FiboSeries(*fFibo))
		}
		if *fPrime > 0 {
			fmt.Printf("the first N prime series of %v is %v", *fPrime, PrimeSeries(*fPrime))
		}
	}
}

//Sum : returns the sum list of integers
func Sum(i []int) int {
	return sum(i)
}

// return sum int of a[] ,using recursion.
func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return sum(a[1:]) + a[0]
}

//Multiply returns the multiply list of integers , by loop
func Multiply(a []int) int {
	res := 1
	for _, v := range a {
		res *= v
	}
	return res
}

// generate the sequence 2, 3, 4 , ....
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// filterPrime : removing those divisible by 'prime'.
func filterPrime(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// PrimeSeries : return prime series
func PrimeSeries(n int) []int {
	var res []int
	ch := make(chan int)
	go generate(ch) // using goroutine to generate sequence.
	for i := 0; i < n; i++ {
		prime := <-ch
		res = append(res, prime)
		chOut := make(chan int)
		go filterPrime(ch, chOut, prime) // using goroutine filter sequance by 'prime' .
		ch = chOut
	}
	return res
}

// generate fibo fibo(n) = fibo(n-1)+fibo(n-1) , using recursion.
func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

// FiboSeries :  return fibonacci series
func FiboSeries(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		res = append(res, fibo(i))
	}
	return res
}
