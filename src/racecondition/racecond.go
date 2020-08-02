package main

import "fmt"
import "time"

var x, y, z int

func mult (){
	fmt.Println(x*y*z)
	x = x+1
	y = y+1
	z = z+1
}

func main(){

	fmt.Println("Ok, this program, tries to multiply some numbers and print their results,")
	fmt.Println("Since there's race conditions, the results can vary")
	fmt.Println("first we got x=0 y=0 z=0, and multiply them with func mult")
	fmt.Println("then add 1 to those variables, so expected result should be:")
	fmt.Println("0*0*0=0 ; 1*1*1=1 ; 2*2*2=8 ; 3*3*3=27")
	fmt.Println("but we have a race condition problem, so if 4 subroutines initializes at the 'same time'")
	fmt.Println("the resuls probably will be 0,0,0,0 ")
	fmt.Println("that happened because every mult function doesn't get the updated x,y,z values")
	fmt.Println("since they starts and gave no time to the variables to update")
	//since i haven't give values to x,y,z their values are 0
	go mult()
	go mult()
	go mult()
	go mult()

	//mult()  //these functions should work properly
	//mult()
	//mult()
	//mult()

	time.Sleep(300 * time.Millisecond) //need to put this or the program doesn't return values
}