import { dirname, join } from "@std/path";

export class solve {
  lines: string[];
  ans: number;
  constructor(input: string, ans = 0) {
    this.lines = input.split("\n");
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

  sortCols(lists: number[][]): number {
    let res = 0;
    const [col1, col2] = lists;
    col1.sort((a, b) => a - b);
    col2.sort((a, b) => a - b);
    for (let idx = 0; idx < col1.length; idx++) {
      const value = Math.abs(col1[idx] - col2[idx]);
      res += value;
    }
    return res;
  }
  part2() {
    const twoCols = this.buildCols(this.lines);
    this.findMatch(twoCols);
  }
  part1() {
    const twoCols = this.buildCols(this.lines);
    const res = this.sortCols(twoCols);
    return res;
  }
  res() {
    return this.ans;
  }
  res1() {
    return this.part1();
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solve(input);
  s1.part1();
  console.log("Part1 result ->", s1.res1());
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}
