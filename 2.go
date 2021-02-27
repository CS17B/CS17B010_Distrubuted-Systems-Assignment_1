package main

import (
	"fmt"
)
const n int = 6
var k int = 0
func Outputmatrix(board[n][n] int){
	    k=k+1
        var i,j int
	for i = 0; i < n; i++{
		for j = 0; j < n; j++{
			fmt.Printf(" %d ", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
func ispossible(board[n][n] int,row int,col int)bool{
	var i,j int
	for i = 0; i < col; i++{
		if board[row][i] == 1{
			return false
		}
	}
	i=row
	j=col
	for ; (i>=0 && j>=0);{
		if board[i][j] == 1{
			return false
		}
		i=i-1
		j=j-1
	}
	i=row
	j=col
	for ; (j>=0 && i<n); {
		if board[i][j] == 1{
			return false
		}
		i=i+1
		j=j-1
	}
	return true
}
func intializeUtil(board[n][n] int,col int)bool{
	if col == n{
		Outputmatrix(board)
		return true
	}

	var res bool = false
	var i int
	for  i = 0; i < n; i++{
		if  ispossible(board, i, col) {
			board[i][col] = 1
			res = intializeUtil(board, col + 1) || res
			board[i][col] = 0
		}
	}
	return res
}

func intialize(){
	board := [n][n]int{}
	if intializeUtil(board, 0) == false{
		fmt.Printf("Solution is not possible")
		return 
	}
	return 
}
func main(){
	intialize()
	fmt.Println("No of Possible arrangements of Queens:",k)
}