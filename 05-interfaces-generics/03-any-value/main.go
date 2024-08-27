package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

const saveFailed = "Save failed."
const saveSuccess = "Save successful."

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printValue(1)
	printValue(1.5)
	printValue("Hello wolrd")
	printValue(false)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}
}

// func printValue(data any) {
func printValue(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("Integer: ", data)
	case float64:
		fmt.Println("Float: ", data)
	case string:
		fmt.Println("String: ", data)
	default:
		fmt.Println("Unsupported data type")
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(saveFailed)
		return err
	}

	fmt.Println(saveSuccess)

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
