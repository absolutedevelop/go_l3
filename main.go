package main 

import "fmt"


func main(){

	fmt.Print("Normal for loop")
	//for loops
	for i:=0; i < 10; i++{
		fmt.Println("current i is", i)
	}

	fmt.Println("Using a counter")
	//using counter
	var counter int = 0
	for counter < 10{
		fmt.Println("Current counter is",counter)

		counter++
	}


}