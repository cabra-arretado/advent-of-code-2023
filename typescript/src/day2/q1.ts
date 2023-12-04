import { readFile } from "../utils";

export const processLine = (line: string): boolean => {
  return true
}

export const processFile = async (input: string): Promise<number> => {
  let total = 0
  for (let ea of input.trim().split('\n')) {
    processLine(ea)
  }
  return total
}

const main = async (): Promise<number> => {
  const input_file = "src/day1/input.txt"
  const data = await readFile(input_file)
  let a = processFile(data)
  return a
}
