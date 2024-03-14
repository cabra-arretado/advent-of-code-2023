import { readFile } from "../utils";

export type Colors = {
  r: number,
  g: number,
  b: number
}

export const processLine = (line: string): number => {
  const arr = line.trim().split(":")
  const day_number = Number(arr[0].split(" ")[1])
  const sprints = arr[1].trim().split(";")
  let clrs: Colors = { r: 0, g: 0, b: 0 }
  for (let ea of sprints) {
    ea = ea.trim()
    for (let i of ea.split(",")) {
      identifyColor(i, clrs)
    }
  }
  return clrs.r * clrs.g * clrs.b
}

export const processFile = async (input: string): Promise<number> => {
  let total = 0
  for (let ea of input.trim().split('\n')) {
    total += processLine(ea)
  }
  return total
}


export const identifyColor = (s: string, clrs: Colors) => {
  const n = Number(s.trim().split(" ")[0])
  if (/red/.test(s) && clrs.r < n) {
    clrs.r = n
  }
  else if (/green/.test(s) && clrs.g < n) {
    clrs.g = n
  }
  else if (/blue/.test(s) && clrs.b < n) {
    clrs.b = n
  }
}

const main = async (): Promise<number> => {
  const data = await readFile("day2")
  let a = processFile(data)
  return a
}

export default main;
