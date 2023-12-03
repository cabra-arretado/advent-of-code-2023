import { readFile } from "../utils";

export const processLine = (line: string): number => {
  let first: string
  const re = /^[0-9]*$/
  for (let i = 0; i < line.length; i++) {
    if (re.test(line[i])) {
      first = line[i]
      break
    }
  }
  let second: string
  for (let i = line.length; 0 <= i; i--) {
    if (re.test(line[i])) {
      second = line[i]
      break
    }
  }
  let n = first.concat(second)
  return Number(n)
}

export const processFile = async (input: string) => {
  const data = await readFile(input)
  let total = 0
  for (let ea of data.trim().split('\n')) {
    total += processLine(ea)
  }
  return total
}

const main = async () => {
  const input = "src/day1/input.txt"
  let a = await processFile(input)
  console.log("Day 1 question 1 answer:", a)
}

main()
