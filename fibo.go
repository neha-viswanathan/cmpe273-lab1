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

/*func main() {
  var length int
  fmt.Println("Enter the length of the fibonacci series")
  fmt.Scanln(&length)
  if(length < 0) {
    fmt.Println("Incorrect/Invalid length for Fibonacci. It cannot be negative!")
  } else {
      var value int = Fibo(length)
      fmt.Println("The returned Fibonacci value is ", value)
  }
}*/
