package main

import ("fmt";"time")

func Sleep(sTime int) {
  select {
  case <-time.After(time.Duration(sTime) * time.Second):
  }
}

func main() {
  fmt.Println("Current time is", time.Now())
  Sleep(5)
  fmt.Println("Current time after 5 sec sleep is")
  fmt.Println(time.Now())
}
