package main

import "fmt"
//import "log"
import "os"
import "bufio"
import	"strings"

// Animal : structure
type Animal struct {
	food, locomotion, noise string
}
//Eat : command for eating
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
//Move : command for movement
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
//Speak : command for sounds
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main(){

	//variables
	var animal Animal

	fmt.Println("welcome, select an animal and an action")
	fmt.Println("each input should be written in the same line, use space")
	fmt.Println("no keyboard input or just one word can 'exit' the program")
	fmt.Println("options of animals are: cow, bird, snake")
	fmt.Println("options of actions are: eat, move, speak")

	for{
		fmt.Println(">")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		name := strings.Split(request, " ")[0]
		info := strings.Split(request, " ")[1]
		name = strings.ToLower(name)
		info = strings.ToLower(info)

		switch name{
		case "cow":
			animal = Animal{"grass","walk","moo"}
		case "bird":
			animal = Animal{"worms","fly","peep"}
		case "snake":
			animal = Animal{"mice","slither","hsss"}
		default:
			fmt.Printf("%s is not a cow, bird or snake", name)
			return	
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("%s is not a valid action", info)
			return
		}

	}

}