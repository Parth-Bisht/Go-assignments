package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Coffee:")

	//comma ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	// fmt.Printf("Type of this rating is %T",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ",numRating+1)
	}
}
