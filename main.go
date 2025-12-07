package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sqweek/dialog"
)

// Colors
const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

// Category Determine Function
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
		fmt.Println("\n" + Red + "--------------- 404 Invalid path ---------------" + Reset)
		fmt.Println(Yellow, "\t \tEnter a valid path \n", err, "\n", Reset)
		footer()
		// return
		os.Exit(1)
	}
}

// footer
func footer() {
	// Center aligned output
	success := Cyan + "Thank you for using File Organizer" + Reset
	developer := Blue + "Developed by Md. Firoz Mahamud" + Reset

	fmt.Println("\n----------------------------------------------------------")
	fmt.Printf("                %s\n", success)
	fmt.Printf("                  %s\n", developer)
	fmt.Println("----------------------------------------------------------")

}

func main() {

	//Take user input from terminal
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter your folder Path: ")
	// scanner.Scan()
	// path := scanner.Text()
	
	// GUI folder picker
	path, err := dialog.Directory().Title("Select Folder to Organize").Browse()
	if err != nil {
		fmt.Println(Yellow + "No folder selected. Exiting..." + Reset)
		return
	}

	fmt.Println(Cyan + "\nSelected folder: " + Reset + path + "\n")

	//reading the directory
	files, err := os.ReadDir(path)
	check(err) //Error Checker

	// count := len(files)//counting the number of files and folders

	// Count only files
	count := 0
	for _, f := range files {
		if !f.IsDir() {
			count++
		}
	}

	//if folder is empty
	if count == 0 {
		fmt.Println("\n" + Yellow + "--------------- Nothing to Organize ---------------" + Reset)
		fmt.Println(Yellow + "\n-------------- This folder is empty ---------------\n" + Reset)
		footer()
		return

	}

	//Displaying the total number of files and folders
	fmt.Println("\n" + Blue + "--------------- Total Number of Files ---------------" + Reset)
	fmt.Println("            ---------------   ", count, "     ---------------      \n")

	//Displaying the files and folders with index
	for _, f := range files {

		//skip all folders
		if f.IsDir() {
			continue
		}
		//skip main organizer folders
		switch f.Name() {
		case "File_Organizer", "Images", "Videos", "Documents", "Audio", "Others":
			continue
		}
		//get the File Extension
		ext := strings.ToLower(filepath.Ext(f.Name()))
		// fmt.Println("[File Extension] :", ext,"\n")
		category := getCategory(ext)

		// First create main folder inside target folder
		mainFolder := filepath.Join(path, "File_Organizer")
		os.MkdirAll(mainFolder, os.ModePerm)

		// fmt.Println("Created:", mainFolder)

		//Create category folder inside file_organizer
		categoryPath := filepath.Join(mainFolder, category)
		os.MkdirAll(categoryPath, os.ModePerm)

		//Move the file to the respective category folder
		oldPath := filepath.Join(path, f.Name())
		newPath := filepath.Join(categoryPath, f.Name())

		err := os.Rename(oldPath, newPath)
		check(err) //Error Checker

		fmt.Printf(Green+"[✓] Moved:"+Reset+" %s → %s\n", f.Name(), category)

	}
	fmt.Println("\n " + Green + "----------- Organizing completed Successfully ------------- " + Reset)
	footer()

	// Pause
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

func init() {
	fmt.Println(Cyan + "----------- WelCome to File Organizer -------------" + Reset)
}
