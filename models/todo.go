package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type ToDoList struct {
	Task      string
	Done      bool
	CreatedAt time.Time
}

type List []ToDoList

func Clean() error {
	file, errs := os.Create("myfile.txt")
	if errs != nil {
		return errs
	}
	defer file.Close()
	return nil
}

func SaveFile(task []ToDoList) {
	file, errs := os.Create("myfile.txt")
	if errs != nil {
		fmt.Println("Error para crear el archivo:", errs)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(task)
	str := string(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	_, errs = file.WriteString(fmt.Sprintf("%+v", str))
	if errs != nil {
		fmt.Println("Error para escribir el archivo:", errs)
		return
	}

}

func ReadFile() (List, error) {
	content, err := os.ReadFile("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	var items []ToDoList
	if err := json.Unmarshal(content, &items); err != nil {
		return []ToDoList{}, err
	}

	return items, nil
}
