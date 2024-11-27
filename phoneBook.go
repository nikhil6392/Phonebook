package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Name 		string 
	Surname 	string
	Tel 		string
	LastAccess  string
}

var CSVFILE = "./csv.data"

var data = []Entry{}

func readCSVFile(filepath string){

}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil

}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	//Updating the index - key does not exist any more
	delete(index, key)
	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func insert(pS *Entry) error {
	// IF Exists , do not insert the entry
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	//Update the index
	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}


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

	fileInfo, err := os.Stat(CSVFILE)
	//Is it a regular file ?

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "Not a regular file !")
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	//Checking the differences between commands
	switch arguments[1] {
	case "insert" :
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-","")  
	

	temp := initS(arguments[2], arguments[3], t)
	//if it was nil, there was an error
	if temp != nil {
		err := insert(temp)
		if err != nil {
			fmt.Println(err)
			return 
	     }
	}

case "delete":
	if len(arguments) != 3 {
		fmt.Println("Usage: delete number")
		return
	}
	t := strings.ReplaceAll(arguments[2], "-", "")
	if !matchTel(t){
		fmt.Println("Not a valid telephone number:", t)
		return
	}
	err := deleteEntry(t)
	if err != nil {
		fmt.Println(err)
	}

case "search":
	if len(arguments) != 3 {
		fmt.Println("Usage: search Number")
		return
	}
	t := strings.ReplaceAll(arguments[2], "-", "")
	if !matchTel(t) {
		fmt.Println("Not a valid telephone number")
		return
	}

	temp := search(t)
	if temp == nil {
		fmt.Println("Number not found:", t)
		return
	}
	fmt.Println(*temp)
case "list":
	list()
default:
	fmt.Println("Not a valid option")
}

}