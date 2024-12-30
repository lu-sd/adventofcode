import { dirname, join } from "@std/path";

export class solve {
  input: string[];
  ans: number;
  constructor(input: string, ans = 0) {
    this.input = input.split("\n");
    this.ans = ans;
  }

  strToInt(s: string): number[] {
    const sList = s.split(" ");
    const numList: number[] = [];
    for (const c of sList) {
      numList.push(Number(c));
    }
    return numList;
  }

  buildLevel(strList: string[]): number[][] {
    const levels: number[][] = [];
    for (const level of strList) {
      const numList = this.strToInt(level);
      levels.push(numList);
    }
    return levels;
  }

  invalid(a: number, b: number): boolean {
    return (a * b <= 0 || Math.abs(b) > 3);
  }

  isSafe(level: number[]): boolean {
    const diff = level[1] - level[0];
    for (let i = 1; i < level.length; i++) {
      const curDif = level[i] - level[i - 1];
      if (this.invalid(diff, curDif)) {
        return false;
      }
    }
    return true;
  }

  isSafe2(level: number[]): boolean {
    let diff = level[1] - level[0];
    let skip = 0;
    for (let i = 1; i < level.length; i++) {
      let curDif = level[i] - level[i - 1];

      if (skip == 1) {
        curDif = level[i] - level[i - 2];
        skip++;
        if (i == 2) {
          diff = level[2] - level[0];
        }
      }

      if (this.invalid(diff, curDif)) {
        if (skip == 0) {
          skip++;
          continue;
        } else {
          return false;
        }
      }
    }
    return true;
  }

  findSafe(levels: number[][]) {
    for (const level of levels) {
      if (this.isSafe(level)) {
        this.ans++;
      }
    }
  }

  part1() {
    const levels = this.buildLevel(this.input);
    this.findSafe(levels);
  }

  findSafe2(levels: number[][]) {
    for (const lev of levels) {
      if (this.isSafe2(lev) || this.isSafe(lev.slice(1))) {
        this.ans++;
      }
    }
  }
  part2() {
    const levels = this.buildLevel(this.input);
    this.findSafe2(levels);
  }
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
  console.log("Part1 result ->", s1.ans);
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.ans);
}
