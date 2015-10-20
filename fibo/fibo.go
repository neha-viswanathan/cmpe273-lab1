package fibo

//Fibonacci sequence
func Fibo(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibo(index-1) + Fibo(index-2)
	}
}
