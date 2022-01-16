package main

import "flag"

func main() {

	// Flag

	//Flag Name , Default Value   , Usage
	timeLimit := flag.Int("limit", 15, "Time limit for Quiz")
	flag.Parse()

	_ = timeLimit

	/* Use of Flag :-
	- It is used to give input to our program from terminal at the time of execution
	- If no input is given it will consider default value as input
	- But It will show 3rd parameter usage when we run it as help
	  like - go run index.go -h (or) -help
	*/
}
