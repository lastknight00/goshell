package dir

import (
	"os"
)

func Dir () ([]os.FileInfo, error) {
	file, err :=os.Open(".")
	
	dir, err := file.Readdir(0)

	
	return dir, err
}

func Cd(path string) error {
	err := os.Chdir(path)
	return err
}