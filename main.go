package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//it'll listen to user input from terminal
	scanner := bufio.NewScanner(os.Stdin)
 	fmt.Print("Please enter your folder Path: ")
	scanner.Scan()
	path := scanner.Text()
	fmt.Println("your folder Path is :", path)
}

func init(){
	fmt.Println("----------- WelCome to Folder Organizer -------------")
}