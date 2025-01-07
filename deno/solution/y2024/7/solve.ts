import { dirname, join } from "@std/path";
import { IntsFromString } from "../../../lib/number.ts";
type oper = (a: number, b: number) => number;

export class solution {
  nums: number[][];
  ans = 0;
  oper1: oper[];
  oper2: oper[];

  constructor(input: string) {
    const lines = input.split("\n");
    const nums: number[][] = [];
    for (const line of lines) {
      const lineToN = IntsFromString(line);
      nums.push(lineToN);
    }
    this.nums = nums;
    this.oper1 = [
      (a: number, b: number) => a + b,
      (a: number, b: number) => a * b,
    ];
    // this.oper2 = [
    //   (a: number, b: number) => a + b,
    //   (a: number, b: number) => a * b,
    //   (a: number, b: number) => {
    //     let fac = 1;
    //     while (fac <= b) {
    //       fac *= 10;
    //     }
    //     return a * fac + b;
    //   },
    //   ]
    this.oper2 = [
      ...this.oper1,
      (a, b) => {
        const fac = 10 ** b.toString().length;
        return a * fac + b;
      },
    ];
  }

  part1() {
    for (const line of this.nums) {
      if (this.isValid(line, this.oper1)) {
        this.ans += line[0];
      }
    }
  }

  isValid(line: number[], oper: oper[]): boolean {
    const target = line[0];
    const input = line.slice(1);
    return this.dfs(1, target, input[0], input, oper);
  }

  dfs(
    level: number,
    target: number,
    curTotal: number,
    input: number[],
    operration: oper[],
  ): boolean {
    if (level === input.length) {
      return target === curTotal;
    }
    if (curTotal > target) {
      return false;
    }
    for (const oper of operration) {
      if (
        this.dfs(
          level + 1,
          target,
          oper(curTotal, input[level]),
          input,
          operration,
        )
      ) {
        return true;
      }
    }
    return false;
  }
  part2() {
    for (const line of this.nums) {
      if (this.isValid(line, this.oper2)) {
        this.ans += line[0];
      }
    }
  }
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
