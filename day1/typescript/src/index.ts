import fs from "fs"

fs.readFile('./src/input.txt', 'utf8', (err, data) => {
  if (err) {
    console.error("Error trying to read file", err)
    return;
  }

  const lines = data.trim().split("\n");

  const list1: number[] = [];
  const list2: number[] = [];

  lines.forEach(line => {
    const parts = line.trim().split(/\s+/)

    if (parts.length === 2) {
      const num1 = Number(parts[0]);
      const num2 = Number(parts[1]);

      if (!isNaN(num1) && !isNaN(num2)) {
        list1.push(num1)
        list2.push(num2)
      }
    }
  })

  list1.sort((a, b) => a - b);
  list2.sort((a, b) => a - b);


  // Iterate over list2 to save on a map

  const repeatCount: Map<number, number> = new Map();

  for (let item of list2) {
    repeatCount.set(item, (repeatCount.get(item) || 0) + 1);
  }

  // Iterate over list 1 and create the similarity score


  let similarityScore = 0;
  let diff = 0;

  for (let i = 0; i < list1.length; i++) {
    if (repeatCount.has(list1[i])) {
      similarityScore += list1[i] * (repeatCount.get(list1[i]) || 0)
    }

    diff += Math.abs(list1[i] - list2[i]);
  }


  console.log("similar score", similarityScore)
  console.log("diff", diff)
})
