import { describe, test, expect } from "@jest/globals"
import { processLine, textToDigit } from "./q2"

describe("Day 1 Question 2", () => {
  test("textToDigit", () => {
    expect(textToDigit("zero")).toBe("zero0zero")
    expect(textToDigit("eightwothree")).toBe("eight8eightwo2twothree3three")
    expect(textToDigit("two1nine")).toBe("two2two1nine9nine")
  })
  test("Test given by the problem", () => {
    const given = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
    const test_cases = given.trim().split("\n")
    const answer = [29, 83, 13, 24, 42, 14, 76]
    for (let i = 0; i < test_cases.length; i++) {
      expect(processLine(test_cases[i])).toBe(answer[i])
    }
  })
})


