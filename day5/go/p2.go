package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Ler o arquivo de entrada
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Dividir o conteúdo do arquivo em duas partes
	parts := strings.Split(string(data), "\n\n")
	depsRaw := strings.TrimSpace(parts[0])
	topRaws := strings.TrimSpace(parts[1])

	// Processar dependências e ordens topológicas
	deps := filterLines(strings.Split(depsRaw, "\n"))
	tops := filterLines(strings.Split(topRaws, "\n"))
	convertedValidTops := []string{}

	// Ajustar e validar cada ordem topológica
	for _, top := range tops {
		orderFiltered := strings.Split(top, ",")
		if !validateOrder(orderFiltered, deps) {
			orderFiltered = adjustOrder(orderFiltered, deps, &convertedValidTops)
		}
	}

	// Calcular o resultado final
	acc := 0
	for _, top := range convertedValidTops {
		nums := strings.Split(top, ",")
		middle := nums[len(nums)/2]
		middleInt, _ := strconv.Atoi(middle)
		acc += middleInt
	}

	fmt.Println("Converted:", convertedValidTops)
	fmt.Println("ValidTops:", acc)
}

// filterLines remove linhas vazias e espaços desnecessários
func filterLines(lines []string) []string {
	var filtered []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			filtered = append(filtered, trimmed)
		}
	}
	return filtered
}

// adjustOrder ajusta a ordem com base nas dependências
func adjustOrder(order []string, deps []string, convertedValidTops *[]string) []string {
	if validateOrder(order, deps) {
		*convertedValidTops = append(*convertedValidTops, strings.Join(order, ","))
		return order
	}

	// Reorganizar a ordem com base nas dependências
	for i := 0; i < len(order); i++ {
		for _, dependence := range deps {
			parts := strings.Split(dependence, "|")
			num1, num2 := parts[0], parts[1]
			idx1 := indexOf(order, num1)
			idx2 := indexOf(order, num2)

			if idx2 != -1 && idx1 != -1 && idx2 < idx1 {
				// Reorganiza os números para respeitar a dependência
				order[idx1], order[idx2] = order[idx2], order[idx1]
			}
		}
	}

	// Recursivamente verifica e ajusta até que a ordem seja válida
	return adjustOrder(order, deps, convertedValidTops)
}

// validateOrder verifica se uma ordem é válida
func validateOrder(order []string, deps []string) bool {
	for j := 0; j < len(order); j++ {
		for _, dependence := range deps {
			parts := strings.Split(dependence, "|")
			num1, num2 := parts[0], parts[1]

			if num2 == order[j] {
				found := false
				var foundIndex int

				for k := 0; k < len(order); k++ {
					if order[k] == num1 {
						found = true
						foundIndex = k
						break
					}
				}

				if found {
					if foundIndex > j {
						return false
					}
				}
			}
		}
	}
	return true
}

// indexOf retorna o índice de um elemento no slice
func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

