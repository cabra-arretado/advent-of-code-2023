import * as fs from "fs"

fs.readFile("src/day1/test.txt", "utf-8", (err, data) => {
  if (err){
    console.error(err)
    return
  }
  console.log(data)
})
