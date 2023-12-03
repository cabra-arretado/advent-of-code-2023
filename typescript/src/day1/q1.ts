import { assert } from "console";
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
  for (let i = line.length; 0 < i; i--) {
    if (/^[0-9]*$/.test(line[i])) {
      second = line[i]
      break
    }
  }
  console.log(Number(first + second))
  return Number(first + second)
}

const processFile = async () => {
  const input = "src/day1/input.txt"
  const data = await readFile(input)
  let total = 0
  for (let ea of data.split('\n')) {
    total += processLine(ea)
  }
  return total
}

// const test = () => {
//   assert(processLine("1abc2") == 12)
//   assert(processLine("pqr3stu8vwx") == 38)
//   assert(processLine("a1b2c3d4e5f") == 15)
//   assert(processLine("treb7uchet") == 77)
// }
