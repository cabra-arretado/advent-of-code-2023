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

export const processFile = async (input: string): Promise<number> => {
  let total = 0
  for (let ea of input.trim().split('\n')) {
    total += processLine(ea)
  }
  return total
}

const main = async (): Promise<number> => {
  const input_file = "src/day1/input.txt"
  const data = await readFile(input_file)
  let a = processFile(data)
  return a
}

export default main
