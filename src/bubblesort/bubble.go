package main

import "fmt"
//import "os"
import "log"
//import "bufio"



func main(){

	//creating  slice for the 10 integers
	Numbers := make([]int,10)
	var temp int
	//var tmp1, tmp2 int //variables for swap
	
	//gathering data for the slice

	fmt.Println("Insert up to 10 integers, one by one or separated with a space")

	for i:=0; i<10; i++{

		 
		if _, err := fmt.Scan(&temp);	err != nil {
			log.Print("  Scan for number failed, due to ", err)
			return
		}


		Numbers[i]=temp
		fmt.Println(Numbers)
	}
	//now that we have all the numbers, lets create the functions
	//func bubblesort() and func swap()
	// need to map Numbers to a [10]int variable
	var Numbers2 [10]int 
	for m:=0;m<10;m++{

		Numbers2[m]=Numbers[m]
	}

	bubblesort(&Numbers2)
	
}

	func swap (x *int, y *int, z *bool){
		if *y > *x{
			fmt.Println("making a swap")
			*x, *y = *y, *x
			*z = true
		}
	}

	func bubblesort (sli *[10]int){

		n:=10
		swaps := true
		for swaps {
			swaps =false
			// cycle runs 10 times, 
			for f:= 1; f < n; f++{
				
				swap(&sli[f] , &sli[f-1], &swaps)
			}
		}
			fmt.Println(sli)
	}