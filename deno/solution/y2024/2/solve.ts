import { dirname, join } from "@std/path";

export class solve {
  input: string;
  lists: string[];
  reports: number[][];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    this.lists = input.split("\n");
    this.reports = this.lists.map((line) => this.strToInt(line));
  }

  strToInt(s: string): number[] {
    const sList = s.split(" ");
    // const numList: number[] = [];
    // for (const c of sList) {
    //   numList.push(Number(c));
    // }
    // return numList;
    return sList.map((v) => Number(v));
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

  part1() {
    for (const report of this.reports) {
      if (this.isSafe(report)) {
        this.ans++;
      }
    }
  }

  part2() {
    for (const report of this.reports) {
      if (this.isSafe(report)) {
        this.ans++;
      } else {
        //skip atmost one item
        for (let skip = 0; skip < report.length; skip++) {
          // const newReport: number[] = [];
          // for (let i = 0; i < report.length; i++) {
          //   if (i == skip) {
          //     continue;
          //   }
          //   const item = report[i];
          //   newReport.push(item);
          // }
          const newReport = report.filter((_, idx) => idx != skip);
          if (this.isSafe(newReport)) {
            this.ans++;
            break;
          }
        }
      }
    }
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
  console.log("Part1 result ->", s1.res());
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.ans);
}
