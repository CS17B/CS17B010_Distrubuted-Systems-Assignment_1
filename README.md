# CS17B010_Distrubuted-Systems-Assignment_1

In Question1 the we want to know the two trees are equivalent we are using two channels and are we are storing values form the tree into channels and
comparing if there are equal or not.Here we are parallely running two channels on two trees to increase speed of excution of code.

In Question2 we want to find all possible arrangements of N Queens in the board we are using backtracking algorithm to solve it it will explore all 
possible combinations and maintains the required conditions like(no two should be in same row and same column and in diagonal).

In Question3 we want to implement counting semaphores using channels. Here in the code I used buffered channels by assigning the size of the buffer that
no of max go routines can only excute at any given point of time. After a delay of 2s some goroutines come out so that new go routines can enter to excute.
