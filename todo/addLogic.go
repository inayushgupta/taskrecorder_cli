package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// add a function to create a dir if not exist
func (list *List) CreateList() {
	filePath := list.Location

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File Does not Exist, creating...")

		dir := filepath.Dir(filePath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}

		file, err := os.Create(filePath)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		fmt.Println("File created at", filePath)
	} else {
		fmt.Println("File already exists")
	}

}

// add task to the file

func (list *List) AddTaskToTheList(task Task) {
	data, err := os.ReadFile(list.Location)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Tasks

	err = json.Unmarshal(data, tasks)
	if err != nil {
		fmt.Println(err)
	}

	tasks = tasks.append(tasks, task)

	newJsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(newJsonData, err)
	}

	err = os.WriteFile(list.Location, newJsonData, 0666)
	if err != nil {
		fmt.Println("err writing to the file", err)
	}
}
