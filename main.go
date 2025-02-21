package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"push/sortMinimum"
)

func hasRepeatingNumbers(arr []int) bool {
	numMap := make(map[int]bool)

	for _, num := range arr {
		if numMap[num] {
			return true
		}
		numMap[num] = true
	}

	return false
}

func main() {
	args := os.Args
	if args[0] == "./push-swap" {
		if len(args) != 2 {
			return
		}
		arr, err := sortMinimum.ToArray(args[1])
		if err != "" {
			fmt.Println("Error")
			return
		}
		if hasRepeatingNumbers(arr) {
			fmt.Println("Error")
			return
		}
		res := sortMinimum.PushMinimum(arr)
		for _, v := range res {
			fmt.Println(v)
		}
	} else if args[0] == "./checker" {
		if len(args) < 2 {
			return
		}
		arr, err := sortMinimum.ToArray(args[1])
		if err != "" {
			fmt.Println("Error")
			return
		}
		if hasRepeatingNumbers(arr) {
			fmt.Println("Error")
			return
		}
		scanner := bufio.NewScanner(os.Stdin)
		var instructions []string
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				break
			}
			instructions = append(instructions, line)
		}

		res := sortMinimum.Checker(arr, instructions)
		fmt.Println(res)
	} else {
		fmt.Println("Error: usage [./push-swap] or [./checker] [arguments]\nExample: ./push-swap 1 2 3 4 5")
	}
}
