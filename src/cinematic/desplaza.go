package main

import "fmt"

func main() {
  var a float64
  var v0 float64
  var s0 float64
  var t float64

  fmt.Print("Enter acceleration:")
  fmt.Scanln(&a)
  
  fmt.Print("Enter the initial velocity:")
  fmt.Scanln(&v0)
  
  fmt.Print("Enter initial displacement:")
  fmt.Scanln(&s0)
  
  fmt.Print("Enter time:")
  fmt.Scanln(&t)

  fn := GenDisplaceFn(a, v0, s0)
  fmt.Println("displacement at time 3: ",fn(3))
	fmt.Println("displacement at time 5: ",fn(5))
	fmt.Println("displacement at time ",t,": ",fn(t))
	
	
}
// GenDisplaceFn : gives insight
func GenDisplaceFn(a,v0,s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return (1/2*a*t*t)+v0*t+s0
	}
}
