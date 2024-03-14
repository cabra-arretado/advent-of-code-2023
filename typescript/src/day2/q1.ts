import { readFile } from "../utils";

export type Colors = {
  r: number,
  g: number,
  b: number
}

const limits: Colors = {
  r: 12,
  g: 13,
  b: 14
}

export const processLine = (line: string): number => {
  const arr = line.trim().split(":")
  const day_number = Number(arr[0].split(" ")[1])
  const sprints = arr[1].trim().split(";")
  for (let ea of sprints) {
    let clrs: Colors = { r: 0, g: 0, b: 0 }
    ea = ea.trim()
    for (let i of ea.split(",")) {
      identifyColor(i, clrs)
      if (limitReached(clrs)){
        return 0
      }
    }
  }
  return day_number
}

export const processFile = async (input: string): Promise<number> => {
  let total = 0
  for (let ea of input.trim().split('\n')) {
    total += processLine(ea)
  }
  return total
}

export const limitReached = (c: Colors): boolean => {
  for (let key of Object.keys(c))
    if (c[key] > limits[key]) {
      return true
    }
  return false
}

export const identifyColor = (s: string, clrs: Colors) => {
  if (/red/.test(s)) {
    clrs.r += Number(s.trim().split(" ")[0])

  }
  else if (/green/.test(s)) {
    clrs.g += Number(s.trim().split(" ")[0])
  }
  else if (/blue/.test(s)) {
    clrs.b += Number(s.trim().split(" ")[0])
  }
}

const main = async (): Promise<number> => {
  const data = await readFile("day2")
  let a = processFile(data)
  return a
}

export default main;
