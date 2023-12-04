import q1 from "./q1"
import q2 from "./q2"

const main = async () => {
  const answer_q1 = await q1()
  const answer_q2 = await q2()
  console.log("~~~~~ Day 1 ~~~~~")
  console.log("Question 1:", answer_q1)
  console.log("Question 2:", answer_q2)
  console.log("~~~~~~~~~~~~~~~~~~")
}
main()