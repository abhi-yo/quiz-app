package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name:")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the quiz!\n", name)

	fmt.Printf("Enter yout age: ")
	var age uint
	fmt.Scan(&age)

	if (age>10) {
		fmt.Println("Yes you can play!")
	}else{
		fmt.Println("You cannot play!")
		return 
	}

	score := 0
	num_questions :=2

	fmt.Println("What is better, the RTX 3090ti or RTX 4080?")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)
	if(answer+ " " +answer2=="RTX 3090ti" || answer + " " + answer2 == "rtx 3090ti"){
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have?")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12{
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n",score,num_questions)
	percent := (float64(score)/float64(num_questions))*100
	fmt.Printf("You scored %v percent.\n",percent)
}