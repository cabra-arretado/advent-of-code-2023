import { assert, warn } from "console";
import { readFile } from "../utils";

// const input = "src/day1/input.txt"
// const data = readFile(input)


const processLine = (line): number => {
  return 0
}

const test = () => {
  assert (processLine("1abc2") == 12)
  assert (processLine("pqr3stu8vwx") == 38)
  assert (processLine("a1b2c3d4e5f") == 15)
  assert (processLine("treb7uchet") == 77)
}
test()
