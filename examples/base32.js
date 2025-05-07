import { b32encode } from "k6/x/workflow_test";

export default function () {
  console.log(b32encode("Hello, World!"))
}
