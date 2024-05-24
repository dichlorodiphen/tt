package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func submitTask(task string) {
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("ACCESS_TOKEN not set in .env file")
	}

	url := "https://api.ticktick.com/open/v1/task"
	client := &http.Client{}

	json := fmt.Sprintf(`{ title: "%s" }`, task)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", body)
		return
	}

	if !(resp.StatusCode == 200 || resp.StatusCode == 201) {
		fmt.Printf("TickTick API responded with an error status: %s\n", body)
		return
	}

	fmt.Println("Task created!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Erroring loading .env file: %v\n", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: tt <task title>")
		return
	}

	task := strings.Join(os.Args[1:], " ")
	submitTask(task)
}
