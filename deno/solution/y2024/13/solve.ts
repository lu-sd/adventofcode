import { dirname, join } from "@std/path";
import { Point } from "../../../lib/dataStructure.ts";
import { IntsFromString } from "../../../lib/number.ts";
type machine = {
  a: Point;
  b: Point;
  target: Point;
};
export class solution {
  input: string;
  lines: string[];
  ans = 0;
  machines: machine[];

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    const machines: machine[] = [];
    const pts: Point[] = [];
    for (const [i, line] of this.lines.entries()) {
      const nums = IntsFromString(line);
      if (i % 4 === 3) {
        machines.push({ a: pts[0], b: pts[1], target: pts[2] });
        continue;
      }
      pts[i % 4] = new Point(0, 0);
      pts[i % 4].x = nums[0];
      pts[i % 4].y = nums[1];
    }
    machines.push({ a: pts[0], b: pts[1], target: pts[2] });
    this.machines = machines;
  }
  part1() {
    for (const m of this.machines) {
      this.ans += this.dfs(m.a, m.b, m.target);
    }
  }

  found(
    aNum: number,
    bNum: number,
    a: Point,
    b: Point,
    target: Point,
  ): boolean {
    return aNum * a.x + bNum * b.x === target.x &&
      aNum * a.y + bNum * b.y === target.y;
  }

  dfs(a: Point, b: Point, target: Point): number {
    let lowest = 0;
    let first = true;
    for (let aNum = 0; aNum < 100; aNum++) {
      for (let bNum = 0; bNum < 100; bNum++) {
        if (this.found(aNum, bNum, a, b, target)) {
          const curL = aNum * 3 + bNum;
          if (first) {
            lowest = curL;
            first = false;
          } else {
            lowest = Math.min(lowest, curL);
          }
        }
      }
    }
    return lowest;
  }
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
