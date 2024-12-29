import { dirname, join } from "@std/path";

export class solution {
  input: string;
  lines: string[];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
  }
  part1() {}
  part2() {}
  res(): number {
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solution(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solution(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}
