package main

import "golang.org/x/tour/tree"
import "fmt"


func Insertinchan(t *tree.Tree, ch chan int) {
	Inordertrav(t, ch)
	close(ch)
}

func Inordertrav(t *tree.Tree, ch chan int) {
    
	if t.Left != nil {
		Inordertrav(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Inordertrav(t.Right, ch)
	}
}


func Equal(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Insertinchan(t1, ch1)
	go Insertinchan(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
    if Equal(tree.New(3), tree.New(2)){
		fmt.Println("Both binary trees are equivalent")
	}else{
	    fmt.Println("Both binary trees are not equivalent")
	}
}