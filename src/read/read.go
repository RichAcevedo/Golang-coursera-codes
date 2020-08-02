package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strings"


func main(){

var filename string
type name struct{
fname string
lname string
}

	names := make([]name,0,3)

	fmt.Println("enter the name of the file, extension included")
	//fmt.Println("the file must be in the same folder")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		filename = scanner.Text()
	}
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan(){
		y := strings.Split(scanner.Text(), " ")
		//if !scanner.Scan(){
		//	break
		//}
	//scanner.Split(bufio.ScanLines)

		var elname name
	elname.fname = y[0]
	elname.lname = y[1]
	names = append(names,elname)
	}
	file.Close()
	
	
	
	for _, eachline:= range names{

		 //fmt.Println(eachline) //testing purpose
		fmt.Println(eachline.fname, eachline.lname)
	}

}