package main

import "bufio"
import "fmt"
import "os"
import "strings"

//Animal : Interface
type Animal interface{
	Eat()
	Move()
	Speak()
}
// Cow : type
type Cow struct{
	food, locomotion, noise string
}
// Bird : type
type Bird struct{
	food, locomotion, noise string
}
// Snake : type
type Snake struct{
	food, locomotion, noise string
}
// functions

//Eat function
// Eat Cow
func (cow Cow) Eat() {
	fmt.Println(cow.food)
}
// Eat Bird
func (bird Bird) Eat(){
	fmt.Println(bird.food)
}
// Eat Snake
func (snake Snake) Eat(){
	fmt.Println(snake.food)
}

//Move function
// Move Cow
func (cow Cow) Move(){
	fmt.Println(cow.locomotion)
}
// Move Bird
func (bird Bird) Move(){
	fmt.Println(bird.locomotion)
}
//Move Snake
func (snake Snake) Move(){
	fmt.Println(snake.locomotion)
}
//Speak function
// Speak Cow
func (cow Cow) Speak(){
	fmt.Println(cow.noise)
}
// Speak Bird
func (bird Bird) Speak(){
	fmt.Println(bird.noise)
}
// Speak Snake
func (snake Snake) Speak(){
	fmt.Println(snake.noise)
}

func main(){

	//making a map
	var NameType map[string]string

	NameType = make(map[string]string)

	//explain how this program works
	fmt.Println("welcome, create a new animal or query for one")
	fmt.Println("each input should be written in the same line, use space")
	//fmt.Println("no keyboard input or just one word can 'exit' the program")
	fmt.Println("commands are: newanimal and query")
	fmt.Println("syntax for new animal is: 'newanimal' 'nameofanimal' 'type of animal' ")
	fmt.Println("types of animals are: cow, bird, snake")
	fmt.Println("syntax for query is: 'query' 'nameofanimal' 'action' ")
	fmt.Println("options for actions are: eat, move, speak")
	fmt.Println("example1: newanimal korin snake")
	fmt.Println("example2: query korin eat")

	//get the data and make some magic

	for{

		fmt.Println("enter the request:")
		fmt.Println(">")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		if len(strings.Split(request," ")) !=3{
			fmt.Println("Commands and arguments should be wrong, input != 3")
			continue
		}

		command := strings.Split(request, " ")[0]
		fmt.Printf("command selected is: %s\n", command)
		
		//check the cases for newanimal and query

		switch command{

		case "newanimal":
			name :=strings.Split(request," ")[1]
			animaltype := strings.Split(request, " ")[2]

			fmt.Printf("name is: %s\n", name)
			fmt.Printf("animal type is: %s\n", animaltype)

				//check if animaltype is cow snake or bird I didn't know how to implement it
			//	if animaltype != "cow"||"snake"||"bird"{
			//		fmt.Printf("%s is not a cow, snake or bird",animaltype)
			//	}

			//adding the names or fail
			_, ok := NameType[name]
			if ok == false{
				NameType[name] = animaltype
				fmt.Println("Create it!")
			} else {
				fmt.Printf("Animal name %s has been taken, try another name\n", name)
			}
			

		case "query":
			name := strings.Split(request, " ")[1]
			info := strings.Split(request, " ")[2]

			fmt.Printf("name is: %s\n", name)
			fmt.Printf("info requested is: %s\n", info)

			//switch for types of animals
			switch NameType[name]{

			case "cow":
				cow := Cow{"grass", "walk", "moo"} //defining slice
				switch info{

				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Printf("%s is not a valid action (eat, move, speak), try again \n",info)
					continue
					
				} //end switch info for cow

			case "bird":
				bird := Bird{"worms", "fly", "peep"} //defining slice
				switch info{

				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Printf("%s is not a valid action (eat, move, speak), try again \n",info)
					continue
					
				} //end switch info for bird

			case "snake":
				snake := Snake{"mice", "slither", "hss"} //defining slice
				switch info{

				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Printf("%s is not a valid action (eat, move, speak), try again \n",info)
					continue
					
				} //end switch info for snake

			default:
				fmt.Printf("%s is not a valid animal (cow, bird, snake), try again\n", NameType[name])


			}//end nametype switch

		default:
			fmt.Printf("%s is not (newanimal, query), try again\n",command)

		} //end switch

	} //end for loop

} //end func main