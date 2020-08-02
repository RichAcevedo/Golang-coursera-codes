package main

import "fmt"
import "os"
import "strconv"

func main(){

	//variable declaration
	var entd string
	
	var sli []int

	fmt.Println("Insert an integer to store in the slice, press 'x' to exit")
	fmt.Scanln(&entd)
		for entd != "x"{

			//inside the loop will make the arrange bigger
			 conv, err := strconv.Atoi(entd)
			 if	err == nil{} 
			sli = append(sli, conv)
			fmt.Println(sli)



			//here the program should see if we exit or continue
			fmt.Println("Insert another integer or press 'x' to exit")
			fmt.Scanln(&entd)
		}
		if entd == "x"{
			os.Exit(0)
		}
	
	//fmt.Println(entd) testing purposes

		
}