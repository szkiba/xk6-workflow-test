import { greeting } from "k6/x/workflow_test";

export default function () {
  console.log(greeting()) // Hello, World!
  console.log(greeting("Joe")) // Hello, Joe!
}
