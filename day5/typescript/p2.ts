import { readFileSync } from 'fs';


const data = readFileSync('input.txt', 'utf8');

const [depsRaw, topRaws] = data.split('\n\n');

const deps: string[] = depsRaw.split(`\n`).filter(line => line.trim() !== '');
const tops: string[] = topRaws.split(`\n`).filter(line => line.trim() !== '');
const validTops: string[] = [];
const convertedValidTops: string[] = [];



const adjustOrder = (order: string[]): string[] => {
  let isValid = validateOrder(order);

  if (isValid) {
    convertedValidTops.push(order.join(","))
    return order; // Retorna a ordem atual se for válida
  }

  // Reorganizar a ordem com base nas dependências
  for (let i = 0; i < order.length; i++) {
    for (const dependence of deps) {
      const [num1, num2] = dependence.split("|");
      const idx1 = order.indexOf(num1);
      const idx2 = order.indexOf(num2);

      if (idx2 !== -1 && idx1 !== -1 && idx2 < idx1) {
        // Reorganiza os números para respeitar a dependência
        [order[idx1], order[idx2]] = [order[idx2], order[idx1]];
      }
    }
  }

  // Recursivamente verifica e ajusta até que a ordem seja válida
  return adjustOrder(order);
};



const validateOrder = (order: string[]) => {
  let viable = true;
  // J is every number of the topological order
  for (let j = 0; j < order.length; j++) {
    for (const dependence of deps) {
      const [num1, num2] = dependence.split("|");
      if (num2 === order[j]) {
        let found = false;
        let foundIndex: number | null = null;

        for (let k = 0; k < order.length; k++) {
          if (order[k] === num1) {
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

  return viable;
}


// Each supposed topological order 
for (const top of tops) {
  const orderFiltered = top.split(",");
  if (!validateOrder(orderFiltered)) {
    adjustOrder(orderFiltered)
  }
}

let acc = 0;

for (const top of convertedValidTops) {
  const nums = top.split(",");
  const middle = nums[Math.floor(nums.length / 2)];
  acc += parseInt(middle)
}

console.log("Converted", convertedValidTops)
console.log("ValidTops", acc)
