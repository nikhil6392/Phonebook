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
	Name string 
	Surname string
	Tel string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search | list <arguments>\n",exe)
		return 
	}
	data = append(data, Entry{"Nikhil","Pathak","384243899"})
	data = append(data, Entry{"Ritik","Pathak","384243299"})
	data = append(data, Entry{"Indra","Pathak","384243199"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry Not Found: ",arguments[2])
			return
		}
		fmt.Println(*result)

	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}

}

func populate(n int, s []Entry){
	for i :=0; i < n; i++{
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100,199))
		data := append(data, Entry{name, surname, n})
	}
}