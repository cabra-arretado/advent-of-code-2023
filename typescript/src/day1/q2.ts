import { readFile } from "../utils";

export const processLine = (line: string): number => {
  line = textToDigit(line)
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

export const processFile = (input: string): number => {
  let total = 0
  for (let ea of input.trim().split('\n')) {
    total += processLine(ea)
  }
  return total
}

export const textToDigit = (line: string): string => {
  const subs: [RegExp, string][] = [
    [/zero/g, "zero0zero"],
    [/one/g, "one1one"],
    [/two/g, "two2two"],
    [/three/g, "three3three"],
    [/four/g, "four4four"],
    [/five/g, "five5five"],
    [/six/g, "six6six"],
    [/seven/g, "seven7seven"],
    [/eight/g, "eight8eight"],
    [/nine/g, "nine9nine"],
  ]
  let s = line
  for (let i = 0; i < subs.length; i++) {
    s = s.replace(subs[i][0], subs[i][1])
  }
  return s
}

const main = async (): Promise<number> => {
  const input_file = "src/day1/input.txt"
  const data = await readFile(input_file)
  let a = processFile(data)
  return a
}

export default main
