package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"math/rand"
	"String"
)

type Entry struct {
	Name 		string 
	Surname 	string
	Tel 		string
	LastAccess  string
}

var CSVFILE = "./csv.data"

var data = []Entry{}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return 
	}
	//If csv file does not exist, create a empty one
	_, err := os.Stat(CSVFILE)
	//If error is not nil , it means that file does not exist
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

}

