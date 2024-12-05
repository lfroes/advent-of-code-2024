
import { readFileSync } from 'fs';
import { argv0 } from 'process';

const input = readFileSync('input.txt', 'utf8');

const matrix = input.split("\n").map((line) => line.split(""));

let countA = 0;


const findOccurrences = (matrix: string[][]): number => {
  const rows = matrix.length;
  const cols = matrix[0].length;
  let occurrences = 0;


  const isOnBounds = (row: number, col: number) =>
    row >= 0 && row < rows && col >= 0 && col < cols;

  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      if (matrix[row][col] !== "A") continue;
      countA++;

      if (
        !isOnBounds(row - 1, col - 1) ||
        !isOnBounds(row + 1, col + 1) ||
        !isOnBounds(row - 1, col + 1) ||
        !isOnBounds(row + 1, col - 1)
      ) {
        continue;
      }

      let diagonalOne = `${matrix[row - 1][col - 1]}${matrix[row + 1][col + 1]}`
      let diagonalTwo = `${matrix[row - 1][col + 1]}${matrix[row + 1][col - 1]}`

      if (
        (
          diagonalOne === "MS" ||
          diagonalOne === "SM"
        ) && (
          diagonalTwo == "MS" ||
          diagonalTwo == "SM"
        )
      ) {
        occurrences++
      }

    }
  }

  return occurrences;
};

const result = findOccurrences(matrix);

console.log("The result is", result, countA);

