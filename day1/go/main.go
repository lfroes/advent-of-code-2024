package main 


import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func main() {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println("Erro ao ler o arquivo", err)
    return
  }


  input := string(data)

  list1, list2 := parseInput(input)

  // creates similarity score 
  similarityScore := calculateSimilarityScore(list1, list2)
  fmt.Println("Similarity Score:", similarityScore)
}

func parseInput(input string) ([]int, []int) {
  var list1, list2 []int
  // Makes each item become a line by itself
  lines := strings.Split(strings.TrimSpace(input), "\n")

  for _, line := range lines {
    parts := strings.Fields(line)

    if len(parts) != 2 {
      continue
    }

    num1, _ := strconv.Atoi(parts[0])
    num2, _ := strconv.Atoi(parts[1])
    list1 = append(list1, num1)
    list2 = append(list2, num2)
  }

  return list1, list2 
}

func calculateSimilarityScore(list1, list2 []int) int {
  // Create a map for list2 occurrencies 

  frequencyMap :=  make(map[int]int)

  for _, num := range list2 {
    frequencyMap[num]++
  }


  similarityScore := 0
  
  for _, num := range list1 {
    if count, exists :=  frequencyMap[num]; exists {
      similarityScore += num * count 
    }
  }

  return similarityScore

}
