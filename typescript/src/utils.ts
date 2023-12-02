import * as fs from "fs"

export const read_file = (file_path: string) => {
  fs.readFile(file_path, "utf-8", (err, data) => {
    if (err) {
      console.error(err)
      return
    }
    return data
  })
}
