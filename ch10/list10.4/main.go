package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getUsedPackages() (string, error) {
	// cmd := exec.Command("go", "list", "-json", ".")
	cmd := exec.Command("go", "list", "-f", `{{join .Deps " "}}`)
	fmt.Printf("args = %s\n", strings.Join(cmd.Args, " "))
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func getAllPackages() (string, error) {
	cmd := exec.Command("go", "list", "-json")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	result, err := getUsedPackages()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
