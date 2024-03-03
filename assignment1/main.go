package main

import (
	"assignment1/entity"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func getArgument(args []string) (int, error) {
	if len(args) > 1 {
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return 0, fmt.Errorf("error converting argument to int: %w", err)
		}
		return arg, nil
	}
	return 0, nil
}

func main() {
	input, err := getArgument(os.Args)
	if err != nil {
		fmt.Println("Error argument")
		return
	}

	student, err := (&entity.Student{}).GetStudent(input)
	if err != nil {
		fmt.Println("Error getting student:", err)
		return
	}

	formattedStudent, err := json.MarshalIndent(student, "", "  ")
	if err != nil {
		fmt.Println("Error formatting:", err)
		return
	}
	fmt.Println(string(formattedStudent))
}
