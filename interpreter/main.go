package main

import "fmt"
import "io/ioutil"
import "os"
import "log"
import "bufio"

func main(){
	if len(os.Args) != 2{
		fmt.Printf("Please specify a brainfuck file")
		os.Exit(1)
	}

	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	b_string := string(bytes)
	s := State{pointer:0, index: 0, lookbackIndex:0}
	interpret(s, b_string)

}

type State struct{
	pointer int
	memory [30000]int64
	index int
	lookbackIndex int
}

func incrementPointer(state *State){
	state.pointer++
}
func decrementPointer(state *State){
	state.pointer--
}
func getMemory(state *State) int64{
	return state.memory[state.pointer]
}
func getLoopback(state *State) int{
	return state.lookbackIndex
}
func incrementIndex(state *State){
	state.index++
}
func setIndex(state *State, i int){
	state.index = i
}
func setLoopback(state *State, i int){
	state.lookbackIndex = i
}
func incrementMemory(state *State){
	state.memory[state.pointer]++
}
func decrementMemory(state *State){
	state.memory[state.pointer]--
}
func checkOpen(state *State) bool{
	if getMemory(state) > 0{
		return true
	}
	return false
}
func interpret(state State, program string){
	reader := bufio.NewReader(os.Stdin)

	for state.index < len(program){
		i := state.index
		current := string(program[i])
		//fmt.Printf(current)
		//fmt.Println("State:", "Pointer:" + strconv.Itoa(state.pointer), "Index:" + strconv.Itoa(state.index), "Memory at pointer:" +strconv.FormatInt(getMemory(&state), 10))
		if current == "+"{
			incrementMemory(&state)
		}
		if current == "-"{
			decrementMemory(&state)
		}
		if current == ">"{
			incrementPointer(&state)
		}
		if current == "<"{
			decrementPointer(&state)
		}
		if current == "["{
			setLoopback(&state, i)
		}
		if current == "]"{
			looping := checkOpen(&state)
			if looping{
				setIndex(&state, state.lookbackIndex)
			}
		}
		if current == "."{
			fmt.Printf(string(rune(getMemory(&state))))
		}
		if current == ","{
			char, err := reader.ReadString('\n')
			fmt.Println(char);
			if err != nil{
				//state.memory[state.pointer] = int(char[0])
			}
		}
		incrementIndex(&state)
	}
}
