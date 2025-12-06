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
	// fmt.Println("your folder Path is :", path)

	//reading the directory
	files, err := os.ReadDir(path)
	count := len(files)

	if err != nil {
	fmt.Println("\n--------------- 404 path not found ---------------")
		fmt.Println("\t Error reading directory path \n", err, "\n")
		return
	}
	fmt.Println("your folder Path is :", path)
	fmt.Println("Type of files  :", fmt.Sprintf("%T", files))
	// fmt.Println("Type of err    :", fmt.Sprintf("%T", err))
	// fmt.Println("Error value    :", err)
	fmt.Println("Total entries of files  :", count)

	

	if count == 0 {
		fmt.Println("This folder is empty.")
		return
	
	}
	fmt.Println("\n--------------- Number of Folder Content ---------------")
	// for _, f := range files {
	// 	if f.IsDir() {
	// 		fmt.Println("[Folder] :", f.Name())
	// 	} else {
	// 		fmt.Println("[File]   :", f.Name())
	// 	}
	// }
	//aded counter for files and folders
	fileCount := 0
	folderCount := 0

	for _, f := range files {
		if f.IsDir() {
			folderCount++
			index := fmt.Sprintf("%02d", folderCount)
			fmt.Printf("[Folder] %s : %s\n", index, f.Name())
		} else {
			fileCount++
			index := fmt.Sprintf("%02d", fileCount)
			fmt.Printf("[File]   %s : %s\n", index, f.Name())
		}
}
}

func init(){
	fmt.Println("----------- WelCome to Folder Organizer -------------")
}