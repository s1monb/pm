/*
Copyright © 2024 SIMON BJØRNØY <SIMON.BJORNOY@GMAIL.COM>
*/
// The project package contains the logic for managing projects.
package project

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/s1monb/proj-mgmt/helpers"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func List() {
	// Read the directory
	files, err := os.ReadDir(viper.GetString("project_directory"))
	if err != nil {
		log.Fatal(err)
	}

	// Create the header for the table
	var table = [][]string{
		{color.WhiteString("NAME"), color.WhiteString("MODIFIED")},
	}

	// Add the files to the table
	for _, file := range files {
		fileInfo, _ := file.Info()
		if !fileInfo.IsDir() {
			continue
		}

		// Format the time and name
		modTime := fileInfo.ModTime().Format("01-02-2006 15:04")
		name, err := url.PathUnescape(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		// Append the file to the table (with colors)
		table = append(table, []string{color.CyanString(name), color.YellowString(modTime)})
	}

	// Print the table
	helpers.PrintTable(table)
}

func New() {
	// Prompt the user for the project name
	prompt := promptui.Prompt{
		Label: "Project name",
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Escape unvalid characters
	validName := url.PathEscape(result)

	// Create the directory
	err = os.Mkdir(viper.GetString("project_directory")+"/"+validName, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Project created.")
}
