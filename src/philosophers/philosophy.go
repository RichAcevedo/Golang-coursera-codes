package main

import "fmt"
import "sync"
import "time"

// ChopS : structure
type chopS struct {sync.Mutex}
// Philo : structure
type Philo struct{
	leftCS, rightCS *chopS
	counter int
}

var ok sync.Once

//starting code
func initial(){
	fmt.Println("Init philosophers")
}

	// eat : function for eating
	func (p Philo) eat (ch chan int){
		ok.Do(initial)
			for i := 0; i < 3; i++{
				<-ch
	
				p.leftCS.Lock()
				p.rightCS.Lock()
	
				fmt.Println("Starting to eat philo ",p.counter)
	
				p.leftCS.Unlock()
				p.rightCS.Unlock()
	
				ch <-i
				fmt.Println("finishing to eat", p.counter)
			}
			//wg.Done()
		return	
	}
	// host : function of the host
	func host (ch chan int, wg *sync.WaitGroup){
		ch <- 1
		ch <- 2
		<- ch

		cSticks := make([]*chopS,5) //making 5 chopsticks

		for i:= 0; i<5; i++{
			cSticks[i] = new(chopS)
		}

		philos := make([]*Philo, 5)	//making philos and associating sticks

	for i:=0; i<5; i++{
		philos[i] = &Philo{cSticks[i], cSticks[(i+1)%5], i+1}
	}

		for i:=0; i<5; i++{ 		//start to eating the philosophers
			go philos[i].eat(ch)
			//wg.Done()			//check
		}
		wg.Done()
	}
	


//MAIN FUNCTION
func main(){

	var wg sync.WaitGroup

	ch := make(chan int, 2) // channel fo capacity 2


	
	wg.Add(1)
	go host(ch, &wg)
	

	wg.Wait()
	time.Sleep(1000*time.Millisecond)
	println("exiting program")



}// end of program


