package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func findOccurrences(matrix [][]rune) (int, int) {
	rows := len(matrix)
	if rows == 0 {
		return 0, 0
	}
	cols := len(matrix[0])
	occurrences := 0
	countA := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] != 'A' {
				continue
			}
			countA++

			if !isOnBounds(row-1, col-1, rows, cols) ||
				!isOnBounds(row+1, col+1, rows, cols) ||
				!isOnBounds(row-1, col+1, rows, cols) ||
				!isOnBounds(row+1, col-1, rows, cols) {
				continue
			}

			diagonalOne := string(matrix[row-1][col-1]) + string(matrix[row+1][col+1])
			diagonalTwo := string(matrix[row-1][col+1]) + string(matrix[row+1][col-1])

			if (diagonalOne == "MS" || diagonalOne == "SM") &&
				(diagonalTwo == "MS" || diagonalTwo == "SM") {
				occurrences++
			}
		}
	}

	return occurrences, countA
}

func main() {
	// Lê o arquivo input.txt
	matrix, err := readMatrix("input.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Procura as ocorrências
	occurrences, countA := findOccurrences(matrix)

	fmt.Printf("O resultado é: %d, A encontrados: %d\n", occurrences, countA)
}

