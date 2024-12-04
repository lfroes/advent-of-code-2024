
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	pattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	scanner := bufio.NewScanner(file)

	isEnabled := true
	solution := 0
	currentIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllStringIndex(line, -1)

		for _, match := range matches {
			absoluteIndex := currentIndex + match[0]
			matchStr := line[match[0]:match[1]]


			switch {

			case matchStr[:3] == "mul":
				if isEnabled {
					numbers := regexp.MustCompile(`\d+`).FindAllString(matchStr, -1)
					n1, _ := strconv.Atoi(numbers[0])
					n2, _ := strconv.Atoi(numbers[1])
					solution += n1 * n2
				}
			case matchStr[:2] == "do" && len(matchStr) == 4:
				isEnabled = true
				fmt.Printf("Enabled at index %d\n", absoluteIndex)
      case matchStr[:5] == "don't":
				isEnabled = false
				fmt.Printf("Disabled at index %d\n", absoluteIndex)
			}
		}

		currentIndex += len(line) + 1
	}

	fmt.Println("Final Solution:", solution)
}

