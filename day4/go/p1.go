package main

import (
	"bufio"
	"fmt"
	"os"
)

var letterValues = map[rune]int{
	'X': 1,
	'M': 2,
	'A': 3,
	'S': 4,
}

func calculateHash(word string) int {
	hash := 0
	for _, char := range word {
		hash = hash*21 + letterValues[char]
	}
	return hash
}

func readMatrix(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return matrix, nil
}

func isOnBounds(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func findOccurrences(matrix [][]rune, targetWord string, targetHash int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	wordLength := len(targetWord)
	occurrences := 0

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 0},
		{0, -1},
		{-1, 1},
		{-1, -1},
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				dRow, dCol := dir[0], dir[1]
				hash := 0
				valid := true

				for i := 0; i < wordLength; i++ {
					newRow := row + dRow*i
					newCol := col + dCol*i

					if !isOnBounds(newRow, newCol, rows, cols) {
						valid = false
						break
					}

					char := matrix[newRow][newCol]
					hash = hash*21 + letterValues[char]
				}

				if valid && hash == targetHash {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func main() {
	// Lê o arquivo input.txt
	matrix, err := readMatrix("input.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Define a palavra-alvo e calcula o hash
	targetWord := "XMAS"
	targetHash := calculateHash(targetWord)

	// Procura a palavra na matriz
	result := findOccurrences(matrix, targetWord, targetHash)

	fmt.Printf("O resultado é: %d\n", result)
}

