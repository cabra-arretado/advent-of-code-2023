import { describe, test, expect } from "@jest/globals"
import { processFile, processLine } from "../day1/q1"

describe("Day 1 Question 1", () => {
  describe("Process Line", () => {
    test("Test given by the problem", () => {
      expect(processLine("1abc2")).toBe(12)
      expect(processLine("pqr3stu8vwx")).toBe(38)
      expect(processLine("a1b2c3d4e5f")).toBe(15)
      expect(processLine("treb7uchet")).toBe(77)
    })
    test("Troublesome lines", () => {
      expect(processLine("4md")).toBe(44)
    })
  })
  describe("ProcessFile", () => {
    test("Mock File", async () => {
      const mock_file = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
      expect(await processFile(mock_file)).toBe(142)
    })
  })
})
