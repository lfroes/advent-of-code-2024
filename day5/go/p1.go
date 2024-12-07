package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	parts := strings.Split(string(data), "\n\n")
	if len(parts) < 2 {
		fmt.Println("Formato inválido do arquivo")
		return
	}

	deps := strings.Split(parts[0], "\n")
	tops := strings.Split(parts[1], "\n")
	var validTops []string

	for _, top := range tops {
		orderFiltered := strings.Split(top, ",")
		viable := true

		for j := 0; j < len(orderFiltered); j++ {
			for _, dep := range deps {
				parts := strings.Split(dep, "|")
				if len(parts) < 2 {
					continue
				}

				num1, num2 := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
				if num2 == orderFiltered[j] {
					found := false
					var foundIndex int

					for k := 0; k < len(orderFiltered); k++ {
						if orderFiltered[k] == num1 {
							found = true
							foundIndex = k
							break
						}
					}

					if found {
						if foundIndex < j {
							viable = true
						}

						if foundIndex > j {
							viable = false
							break
						}
					}
				}
			}
			if !viable {
				break
			}
		}

		if viable {
			validTops = append(validTops, top)
		}
	}

	acc := 0
	for _, top := range validTops {
		nums := strings.Split(top, ",")
		middle := nums[len(nums)/2]
		middleValue, err := strconv.Atoi(strings.TrimSpace(middle))
		if err != nil {
			fmt.Println("Erro ao converter número:", err)
			continue
		}
		acc += middleValue
	}

	fmt.Println("ValidTops:", acc)
}

