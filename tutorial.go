package main

import "fmt"

func main() {
	fmt.Println("welcome to my quiz game")
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("yOU CAN PLAY")
	} else {
		fmt.Println("YOU CANNOT play")
		return
	}
	score := 0
	fmt.Println("please give answer is capital letters")
	fmt.Printf("WHO is the present president of india (2023) :")
	var answer string
	var answer1 string
	fmt.Scan(&answer, &answer1)
	fmt.Println(answer, answer1)
	if answer+" "+answer1 == "DROUPADI MURMU" {
		fmt.Println("correct answer")
		score++
	} else {
		fmt.Println("wrong answer")

	}
	fmt.Printf("HOW MANY NUMBER OF D DO WORD DDPS HAVE  :")
	var num int
	fmt.Scan(&num)
	if num == 2 {
		fmt.Println("correct answer")
		score++
	} else {
		fmt.Println("wrong answer")

	}
	fmt.Printf("Your final score is %v", score)
}
