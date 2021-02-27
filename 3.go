package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)
var ch = make(chan int, 4) 

func main(){
  wg := sync.WaitGroup{}
  arr := [10]int{5,10,15,20,25,30,35,40,45,50}
  for _,n:= range arr {
    wg.Add(1)
    go Printval(&wg, n)
  }
  wg.Wait()
}

func Printval(wg* sync.WaitGroup, n int) {
  var e int  
  e =rand.Intn(10)
  ch <- e
  fmt.Println(n)
  time.Sleep(2000 * time.Millisecond)
  <- ch 
   wg.Done()
}
