package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//Category Determine Function
func getCategory(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff":	
		return "Image"	
	case ".mp4", ".mkv", ".avi", ".mov", ".wmv":
		return "Video"
	case ".mp3", ".wav", ".ogg", ".flac":
		return "Audio"
	case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx":
		return "Document"	
	default:
		return "Other"
	}
}

// Error Checker
func check(err error) {
	if err != nil {
		fmt.Println("\n--------------- 404 invalid path ---------------")
		fmt.Println("\t \tEnter a valid path \n", err, "\n")
		// return
		os.Exit(1)
	}
}


func main() {

	//Take user input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your folder Path: ")
	scanner.Scan()
	path := scanner.Text()

	//reading the directory
	files, err := os.ReadDir(path)
	check(err) //Error Checker

	count := len(files)//counting the number of files and folders

	//if folder is empty
	if count == 0 {
		fmt.Println("This folder is empty.")
		return

	}

	//Displaying the total number of files and folders
	fmt.Println("\n--------------- Total Number of Folder Content ---------------")
	fmt.Println("            ---------------   ", count, "     ---------------      \n")

	//Displaying the files and folders with index
	fileCount := 0
	folderCount := 0

	for _, f := range files {
		if f.IsDir() {
			folderCount++
			// index := fmt.Sprintf("%02d", folderCount)
			// fmt.Printf("[Folder]      %s : %s\n", index, f.Name())
		} else {
			fileCount++
			// index := fmt.Sprintf("%02d", fileCount)
			// fmt.Printf("[File]        %s : %s\n", index, f.Name())
		}

		//get the File Extension
		ext := strings.ToLower(filepath.Ext(f.Name()))
		// fmt.Println("[File Extension] :", ext,"\n")
		category := getCategory(ext)

		//Create Category Folder if not exists
		categoryPath := filepath.Join(path, category)
		os.MkdirAll(categoryPath, os.ModePerm)

		//Move the file to the respective category folder
		oldPath := filepath.Join(path, f.Name())
		newPath := filepath.Join(categoryPath, f.Name())

		err := os.Rename(oldPath, newPath)
		check(err) //Error Checker

		fmt.Println("Moved:", f.Name(), "→", category)

	}

		fmt.Println("\nOrganizing complete!")
	
}

func init() {
	fmt.Println("----------- WelCome to Folder Organizer -------------")
}
