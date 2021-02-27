package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)
var tokens = make(chan int, 4) 

func main(){
  wg := sync.WaitGroup{}
  arr := [10]int{5,10,15,20,25,30,35,40,45,50}
  for _,n:= range arr {
    wg.Add(1)
    go doInParallel(&wg, n)
  }
  wg.Wait()
}

func doInParallel(wg* sync.WaitGroup, n int) {
  var e int  
  e =rand.Intn(10)
  tokens <- e
  fmt.Println(n)
  time.Sleep(2000 * time.Millisecond)
  <- tokens 
   wg.Done()
}