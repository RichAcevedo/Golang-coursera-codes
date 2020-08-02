package main

import "fmt"
import "bufio"
import "encoding/json"
import "os"

func main(){

	//initialize variables 
	var temp1,temp2 string
	//idMap := map[string]string{
	//	"juan":"123 st"}
	

	//data entry

	fmt.Println("enter the name of the person")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		temp1 = scanner.Text()
	}
	fmt.Println("enter the address of the person")
	scq := bufio.NewScanner(os.Stdin)
	scq.Scan()
	if scq.Err() == nil {
		temp2 = scq.Text()
	}

		//creating a map
		var match = make(map[string]string)
		
		match["name"] = temp1
		match["address"] = temp2
		//just to check evrything is ok
		fmt.Println("map in golang: ",match)

		//import to JSON
		//i needed to make a struct
		/*type argonaut struct{
			names string
			addresses string
		}
		var p1 argonaut
		p1.names =temp1	
		p1.addresses=temp2*/

		//fmt.Println(p1)
		barr,err:=json.Marshal(match)
		if err == nil{}
			fmt.Println(string(barr))
		


}