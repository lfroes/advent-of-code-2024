import { readFileSync } from "fs"


function isSafe(numbers: Array<number>) {
  let isCrescent = true;
  let isDecrescent = true;


  for (let i = 0; i < numbers.length - 1; i++) {
    let diff = Math.abs(numbers[i + 1] - numbers[i])

    if (diff > 3 || diff < 1) {
      return false
    }

    if (numbers[i + 1] < numbers[i]) {
      isCrescent = false
    }

    if (numbers[i + 1] > numbers[i]) {
      isDecrescent = false
    }
  }

  return isCrescent || isDecrescent
}


const solution = (path: string) => {
  let safeList = 0;
  try {
    const content = readFileSync(path, 'utf-8')
    const lines = content.split("\n");

    lines.forEach((line, index) => {
      const strNumbers = line.split(" ").filter((num) => num.trim() !== "");
      const numbers = strNumbers.map((num) => Number(num));

      if (numbers.length === 0) {
        return;
      }

      if (isSafe(numbers)) {
        safeList++
      } else {
        for (let i = 0; i < numbers.length; i++) {
          let temp = [...numbers.slice(0, i), ...numbers.slice(i + 1)]
          if (isSafe(temp)) {
            safeList++;
            break;
          }
        }
      }
    })
  } catch (err) {
    console.error("An error has ocurred", err)
  }

  console.log("Safe List:", safeList)
}


solution("./input.txt")
