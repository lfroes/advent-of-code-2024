import { readFileSync } from 'fs';


const data = readFileSync('input.txt', 'utf8');

const [depsRaw, topRaws] = data.split('\n\n');

const deps: string[] = depsRaw.split(`\n`).filter(line => line.trim() !== '');
const tops: string[] = topRaws.split(`\n`).filter(line => line.trim() !== '');
const validTops: string[] = [];


// Each supposed topological order 
for (const top of tops) {
  const orderFiltered = top.split(",");
  // J is every number of the topological order
  let viable = true;
  for (let j = 0; j < orderFiltered.length; j++) {
    // iterate over each node to find if has any Dependencies
    let stop = false;
    for (const dependence of deps) {
      const [num1, num2] = dependence.split("|");
      if (num2 === orderFiltered[j]) {
        // The number has a dependence
        let found = false;
        let foundIndex: number | null = null;

        for (let k = 0; k < orderFiltered.length; k++) {
          if (orderFiltered[k] === num1) {
            found = true
            foundIndex = k;
            break
          }
        }

        if (found && foundIndex !== null) {
          if (foundIndex < j) {
            viable = true;
          }

          if (foundIndex > j) {
            viable = false
            break;
          }
        }
      }
    }

    if (!viable) break;
  }

  if (viable) {
    validTops.push(top)
  }
}

let acc = 0;

for (const top of validTops) {
  const nums = top.split(",");
  const middle = nums[Math.floor(nums.length / 2)];
  acc += parseInt(middle)
}

console.log("ValidTops", acc)
