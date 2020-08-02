package main

import "fmt"
import "log"



func main(){
	
	var acc, invo, indis, time float64

	//ask for a, vo, so
	
	fmt.Println("Insert acceleration a: ")	
		if _, err := fmt.Scan(&acc);	err != nil {
			log.Print("  Scan for number failed, due to ", err)
			return
		}
	fmt.Println("Insert initial velocity Vo: ")	
		if _, err := fmt.Scan(&invo);	err != nil {
			log.Print("  Scan for number failed, due to ", err)
			return
		}
	fmt.Println("Insert initial displacement So: ")	
		if _, err := fmt.Scan(&indis);	err != nil {
			log.Print("  Scan for number failed, due to ", err)
			return
		}
	fmt.Println("Insert evaluation time to: ")	
		if _, err := fmt.Scan(&time);	err != nil {
			log.Print("  Scan for number failed, due to ", err)
			return
		}
	 fmt.Println(acc,invo,indis,time) //only used for debugging

	 //finalfunc := GenDisplaceFn(acc,invo,indis)
	//fmt.Println(finalfunc(time))
		//first function to compute acceleration, initial velocity
		//and initial displacement
		//fn will give us the total displacement in function of time
		
		// GenDisplaceFn : gives us the displacement
func GenDisplaceFn(a float64,Vo float64,So float64) 
func (float64) float64{

fn:= func (t float64) float64{
	return (0.5*a*(t**2.0))+(Vo*t)+(So) 
}
return fn
}

}

