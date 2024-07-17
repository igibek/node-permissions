package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var path string
var builtins []string

func detectLibrary(path string) []string {
	return nil
}
func traverse(path string) {

}

func main() {
	fmt.Println("node permissions tool")
	flag.StringVar(&path, "path", ".", "Path to the Node application source folder")
	flag.Parse()
	if path == "." {
		p, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		path = p
	}
	fmt.Print(path)

	// change the current working directory
	os.Chdir(path)

}
