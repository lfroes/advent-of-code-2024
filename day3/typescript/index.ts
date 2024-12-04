import { readFileSync } from 'fs';

const scanMemory = (path: string) => {
  const content: string = readFileSync(path, 'utf-8');

  const matches = content.match(/mul\(\d+,\d+\)|do\(\)|don't\(\)/g);

  console.log(matches);

  let acc = 0;
  let isEnabled = true;

  if (matches) {
    for (let i = 0; i < matches.length; i++) {
      if (matches[i].includes("don't")) {
        isEnabled = false;
        console.log("dont", i)
        continue;
      }

      if (matches[i].includes('do')) {
        isEnabled = true;
        console.log("do", i)
        continue;
      }

      if (isEnabled && matches[i].includes('mul')) {
        const [a, b] = matches[i]
          .replace('mul(', '')
          .replace(')', '')
          .split(',')
          .map((x) => parseInt(x));

        acc += a * b;
      }
    }
  }

  console.log(matches?.length)

  console.log(acc);
}

scanMemory('input.txt');
