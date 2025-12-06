package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//it'll listen to user input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your folder Path: ")
	scanner.Scan()
	path := scanner.Text()
	// fmt.Println("your folder Path is :", path)

	//reading the directory
	files, err := os.ReadDir(path)
	count := len(files)

	if err != nil {
		fmt.Println("\n--------------- 404 invalid path ---------------")
		fmt.Println("\t \tEnter a valid path \n", err, "\n")
		return
	}

	if count == 0 {
		fmt.Println("This folder is empty.")
		return

	}
	fmt.Println("\n--------------- Total Number of Folder Content ---------------")
	fmt.Println("            ---------------   ", count, "     ---------------      \n")

	//aded counter for files and folders
	fileCount := 0
	folderCount := 0

	for _, f := range files {
		if f.IsDir() {
			folderCount++
			index := fmt.Sprintf("%02d", folderCount)
			fmt.Printf("[Folder]      %s : %s\n", index, f.Name())
		} else {
			fileCount++
			index := fmt.Sprintf("%02d", fileCount)
			fmt.Printf("[File]        %s : %s\n", index, f.Name())
		}

		//detect the File Extension
		ext := strings.ToLower(filepath.Ext(f.Name()))
		fmt.Println("[File Extension] :", ext,"\n")

	}

	
}

func init() {
	fmt.Println("----------- WelCome to Folder Organizer -------------")
}
