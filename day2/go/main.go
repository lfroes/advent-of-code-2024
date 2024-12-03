package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func AbsInt(n int) int {
  if n < 0 {
    return -n 
  } 

  return n 
}

func IsSafe(numbers []int) bool {
  isCrescent := true
  isDecrescent := true 

  for i := 0; i < len(numbers) - 1; i++ {
    diff := AbsInt(numbers[i + 1] - numbers[i])

    if diff > 3 || diff < 1 {
      return false
    }

    if numbers[i+1] > numbers[i] {
      isDecrescent = false
    } else if numbers[i+1] < numbers[i] {
      isCrescent = false
    }
  }

  return isCrescent || isDecrescent
}

func RemoveAtIndex(slice []int, index int) []int {
  newSlice := make([]int, len(slice) -1)
  copy(newSlice, slice[:index])
  copy(newSlice[index:], slice[index+1:])
  return newSlice
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file", err)
    return
  }
  defer file.Close()

  safeLists := 0
  scanner := bufio.NewScanner(file)
  
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Fields(line)
    numbers := make([]int, len(parts))

    for i, part := range parts {
      num, _ := strconv.Atoi(part)
      numbers[i] = num
    }

    if IsSafe(numbers) {
      safeLists++
    } else {
      for i := 0; i < len(numbers); i++ {
        temp := RemoveAtIndex(numbers, i)
        if IsSafe(temp) {
          safeLists++
          break
        }
      }
    }
  }


  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading the file", err)
    return
  }

  fmt.Println("Safe Lists", safeLists)
}

