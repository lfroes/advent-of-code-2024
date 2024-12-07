import { readFileSync } from 'fs';


const data = readFileSync('input.txt', 'utf8');

// Let`s try use DFS to solve this problem
// We need to create a graph from the input

const lines: string[][] = data.split('\n').filter(line => line.trim() !== '')
  .map(line => line.split(""))


// Create a map with the directions, and the next direction to go case needed
const directions: Map<string, string> = new Map();
directions.set("^", ">");
directions.set(">", "v");
directions.set("v", "<");
directions.set("<", "^");


// Find the initial State position of the guard
let currentPos: number[] = [0, 0] // X, Y
let currentDirection: string = "^";
let isInside: boolean = true;
for (let i = 0; i < lines.length; i++) {
  let pos = lines[i].indexOf(currentDirection)
  if (pos !== -1) {
    console.log("Starting point found", pos, i);
    currentPos = [i, pos]
  }
}


// We need now to move the guard through the map, and saves every postion that he has visited
const visitedPos: Set<string> = new Set();
visitedPos.add(currentPos.join(","));


// Checks if the guard did not leave map boundrys
const isInsideTheZone = (pos: number[]) => {
  let zoneX = lines.length;
  let zoneY = lines[0].length;
  let [x, y] = pos;
  let check = x >= 0 && x < zoneX && y >= 0 && y < zoneY;

  if (!check) {
    currentPos = pos;
  }

  return check;
}

// Move the guard
const move = (pos: number[], direction: string) => {
  // console.log(directions.has(direction))
  if (!directions.has(direction)) {
    console.log("ERR: direction not found in directions");
    // safety measure to kill while loop
    return currentPos = [-9999, -9999]
  }

  let dir = [...directions.keys()].find(key => key === direction);
  let nextDir: string = directions.get(direction) as string;

  let nextZone: number[] = [0, 0]

  switch (dir) {
    case "v":
      nextZone = [pos[0] + 1, pos[1]];
      break;
    case ">":
      nextZone = [pos[0], pos[1] + 1];
      break;
    case "^":
      nextZone = [pos[0] - 1, pos[1]]
      break;
    case "<":
      nextZone = [pos[0], pos[1] - 1];
      break;
  }

  if (
    nextZone[0] < 0 ||
    nextZone[0] >= lines.length ||
    nextZone[1] < 0 ||
    nextZone[1] >= lines[0].length
  ) {
    console.log("Out of bounds:", nextZone);
    return (currentPos = [-9999, -9999]); // Finaliza o movimento
  }

  let area = lines[nextZone[0]][nextZone[1]];

  // Check if next zone is not obstructed
  if (area === "#") {
    // Rotate 90 degress
    return currentDirection = nextDir;
  }

  // Check if is not already visited 

  if (!visitedPos.has(nextZone.join(","))) {
    visitedPos.add(nextZone.join(","))
  }

  // Updates current position
  return currentPos = nextZone;
}


// move(currentPos, currentDirection)
while (isInsideTheZone(currentPos)) {
  move(currentPos, currentDirection)
}


console.log("Finished moving around, visited:", visitedPos.size, "positions!");



