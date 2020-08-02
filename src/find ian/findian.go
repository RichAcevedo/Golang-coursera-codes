package main

import "fmt"
import "strings"
import "bufio"
import "os"
//import "unicode"

func main() {
	var ian string


	fmt.Printf("Insert a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		ian = scanner.Text()
	}
	
	//fmt.Scan(&ian)
	//modify the string, all in lowercase
	trueian := strings.ToLower(ian)
	//trueian = strings.Replace(trueian," ","",-1)
	//fmt.Println(trueian)
	fmt.Println(trueian)
	//check if string starts with i
	if strings.HasPrefix(trueian,"i")==true{
		//now check if string finishes with n

		if strings.HasSuffix(trueian,"n")==true {
			//if we are here, only we need to find "a"

			if strings.Contains(trueian,"a")== true{
			//found it!
			fmt.Println("Found!")
			}else {
				fmt.Printf("Not Found")
			}
		}else {
			fmt.Printf("Not Found")
		}
	
		//fmt.Printf("Found!") testing used
	}else {
		fmt.Printf("Not Found")
	}

	
}
