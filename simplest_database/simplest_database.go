package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide arguments.")
		return
	}

	var command string = os.Args[1]
	var key string = os.Args[2]

	file, err := os.OpenFile("data.db", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if command == "set" {
		var value string = os.Args[3]
		set(key, value, file)

		fmt.Println("Successfully set", key, "to", value)
	} else if command == "get" {
		value, _ := get(key, file)

		fmt.Println("Value for", key, "is", value)
	}
}

func set(key string, value string, file *os.File) {
	fmt.Println("Setting", key, "to", value)

	_, err := file.WriteString(key + "," + value + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func get(key string, file *os.File) (string, error) {
	var lastValue string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var k, v string
		parts := strings.Split(line, ",")
		k, v = parts[0], parts[1]
		if k == key {
			lastValue = v
		}
	}

	if lastValue == "" {
		fmt.Println("No value found for", key)
		return "", nil
	}

	return lastValue, nil
}
