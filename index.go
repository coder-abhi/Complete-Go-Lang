package main

func main() {

	/*

		//Flag

		// Use of Flag :-
		// - It is used to give input to our program from terminal at the time of execution
		// - If no input is given it will consider default value as input
		// - But It will show 3rd parameter usage when we run it as help
		//   like - go run index.go -h (or) -help


						//Flag Name , Default Value   , Usage
		timeLimit := flag.Int("limit", 15, "Time limit for Quiz")
		flag.Parse()

		_ = timeLimit //to avoid error


		// Input in terminal to run file :- - go run index.go -limit=30
		// Output :- (value of timeLimit varible is 30 Now in Program)

	*/

	/*

		//Generating Random Numbers winthin range

		min := 0   //range - lowerbound
		max := 100 //range - upperbound

		rand.Seed(time.Now().UnixNano())
		digit := rand.Intn(max-min+1) + min // Now Everytime Code run digit willl get different No.

		fmt.Println(digit)

		// Input Run Program 5 times
		// it will print different numbers 5 times

	*/
}
