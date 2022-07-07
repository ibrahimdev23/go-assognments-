package main


//1. sum is a function which takes a slice of numbers and adds them together.
// What would its function signature look like in Go?

sum(numbers []int) int

//Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
// For example half(1) should return (0, false) and half(2) should return (1, true).

func halves(num int) (int, bool){
	if(num % 2 == 0){
		return num / 2, true
	} else {
		return num / 2, false 
	}
}



//3). Write a function with one variadic parameter that finds the
// greatest number in a list of numbers.

func findGreatest(args ...int) int {
	highest := args[0]
	for _, i := range args {
		if i > highest {
			highest = i 
		}

	}
	return highest 
}


//4)Using makeEvenGenerator as an example, 
//write a makeOddGenerator function that generates odd numbers.

func makeOddGenerator() func() int {
	i := 1
	return func() (oddNum int){
		oddNum = i 
		i += 2
		return
	}
}



//5)The Fibonacci sequence is defined as: 0, fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). 
//Write a recursive function which can find fib(n)

func fib(n int) int{
	if(n == 0){
		return 0
	} else if (n == 1){
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}