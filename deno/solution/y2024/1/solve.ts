import { dirname, join } from "@std/path";

export class solve {
  input: string;
  ans: number;
  constructor(input: string, ans = 0) {
    this.input = input;
    this.ans = ans;
  }

  buildCols(lines: string[]): number[][] {
    const col1: number[] = [];
    const col2: number[] = [];
    for (const line of lines) {
      const [n1, n2] = line.split("   ");
      col1.push(Number(n1));
      col2.push(Number(n2));
    }
    return [col1, col2];
  }

  findMatch(lists: number[][]) {
    const [col1, col2] = lists;
    const col2Freq = new Map<number, number>();

    for (const n2 of col2) {
      col2Freq.set(n2, (col2Freq.get(n2) || 0) + 1);
    }
    for (const n1 of col1) {
      this.ans += n1 * (col2Freq.get(n1) || 0);
    }
  }
  part1() {
    const lines = this.input.split("\n");
    const twoCols = this.buildCols(lines);
    this.findMatch(twoCols);
  }
  part2() {}
  res() {
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solve(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}
