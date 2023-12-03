import { readFile } from "../utils";

const processLine = (line: string): number => {
  let first: string
  for (let i = 0; i < line.length; i++) {
    if (/^[0-9]*$/.test(line[i])) {
      first = line[i]
      break
    }
  }
  let second: string
  for (let i = line.length; 0 <= i; i--) {
    if (/^[0-9]*$/.test(line[i])) {
      second = line[i]
      break
    }
  }
  let n = first.concat(second)
  return Number(n)
}

const processFile = async (input: string) => {
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
  console.log(a)
}

main()
