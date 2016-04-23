package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
	"dir"
)

func main() {
	var args []string
	reader := bufio.NewReader(os.Stdin)
	
	args = inputCommand(reader)
	for args[0] != "exit" && args[0] != "quit" {
		
		switch args[0] {
		case "cd":
			err := dir.Cd(args[1])
			if err != nil {
				fmt.Print("error\n")
			}
		case "dir":
			fileList, err := dir.Dir()
			if err != nil {
				fmt.Printf("%s\n", err) 
			} else {
				showDirList(fileList)
			} 
		}
		args = inputCommand(reader)
	}
}

func inputCommand(reader *bufio.Reader) []string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("error\n")
	}
	fmt.Printf("%s>", wd)
	input, _ := reader.ReadString('\n')
	retval := strings.Split(input, " ")
	for index, _ := range retval {
		retval[index] = strings.TrimSpace(retval[index])
	}
	return retval
}

func showDirList(fileList []os.FileInfo) {
	var isDir string
	for _, a := range fileList {
		if a.IsDir() {
			isDir = "Dir"
		} else {
			isDir = "File"
		} 
		fmt.Printf("%d\t%s\t%s\t%s\n", a.Size(), a.Mode(), isDir, a.Name())
	}
}