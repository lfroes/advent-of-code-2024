import { readFileSync } from 'fs';

const input = readFileSync('input.txt', 'utf8');

const letterValues: Record<string, number> = { X: 1, M: 2, A: 3, S: 4 };

const calculateHash = (word: string): number => {
  return word.split("").reduce((hash, char) => hash * 21 + (letterValues[char] || 0), 0);
};

const targetWord = "XMAS";
const targetHash = calculateHash(targetWord);

const matrix = input.split("\n").map((line) => line.split(""));

const findOccurrences = (matrix: string[][], targetWord: string, targetHash: number) => {
  const rows = matrix.length;
  const cols = matrix[0].length;
  const wordLength = targetWord.length;
  let occurrences = 0;

  const directions = [
    [0, 1],  // start from left to r 
    [1, 0],  // top bottom
    [1, 1],  // bottom right
    [1, -1], // bottom left
    [-1, 0], // top to bottom
    [0, -1], // right to left
    [-1, 1], // top right
    [-1, -1] // top left
  ]

  const isOnBounds = (row: number, col: number) =>
    row >= 0 && row < rows && col >= 0 && col < cols

  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      for (const [dRow, dCol] of directions) {
        let hash = 0;
        let valid = true;

        for (let i = 0; i < wordLength; i++) {
          const newRow = row + dRow * i;
          const newCol = col + dCol * i;

          if (!isOnBounds(newRow, newCol)) {
            valid = false;
            break;
          }

          const char = matrix[newRow][newCol];
          hash = hash * 21 + (letterValues[char] || 0);
        }

        if (valid && hash === targetHash) {
          occurrences++;
        }
      }
    }
  }

  return occurrences
}

const result = findOccurrences(matrix, targetWord, targetHash)

console.log("The result is", result)
