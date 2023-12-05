import * as fs from "fs"

export const readFile = async (file_path: string) => {
  return fs.readFileSync(file_path, "utf-8")
}

const printInColors = (color: string, s: any) => {
  switch (color) {
    case "red":
      console.log(`\x1b[31m${s}\x1b[0m`)
      break
    case "blue":
      console.log(`\x1b[34m${s}\x1b[0m`)
      break
    case "yellow":
      console.log(`\x1b[33m${s}\x1b[0m`)
  }
}

const blank = () => "To be done"


export const printAnswers = async (day: string | number, q1: () => any = blank, q2: () => any = blank) => {
  const answer_q1 = await q1()
  const answer_q2 = await q2()
  printInColors("blue", `~~~~~ Day ${day} ~~~~~`)
  printInColors("red", "Question 1:")
  printInColors("yellow", answer_q1)
  printInColors("red", "Question 2:")
  printInColors("yellow", answer_q2)
 printInColors("blue", "~~~~~~~~~~~~~~~~~")
}

