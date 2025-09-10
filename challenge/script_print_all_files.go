package main
import (
		"fmt"
    "os"
	  
)
 
func main() {
	root:= "."
	// filepath.Walk(root, func(path string, info fs.FileInfo, err error) error 
	spisok, err := os.ReadDir(root)
	if err != nil {
	fmt.Println(err)

	}
	fmt.Println(spisok)
}
