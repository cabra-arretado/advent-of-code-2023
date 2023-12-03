import * as fs from "fs"

export const readFile = async (file_path: string) => {
  return fs.readFileSync(file_path, "utf-8")
}
