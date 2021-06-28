package main 

import (
	"os"
	"fmt"

)

func main() {
	// Open the file
	file, err := os.Open("test.txt")
	// Handles the error if their is any
	if err != nil {
		return
	}
	// Close the file
	defer file.Close()

	// Get the size of the file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// Read the file
	bs := make([]byte, stat.Size())

	_, err := file.Read(bs)

	if err != nil {
		return
	}

	str := string(bs)

	fmt.Println(str)


}