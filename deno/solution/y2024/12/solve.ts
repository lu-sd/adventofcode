import { dirname, join } from "@std/path";
import { Dirs4, Grid, Point } from "../../../lib/dataStructure.ts";

export class solution {
  input: string;
  ans = 0;
  grid: Grid<string>;
  land = new Map<string, Point[]>();
  seen = new Set<string>();

  constructor(input: string) {
    this.input = input;
    const lines = input.split("\n");
    const nrow = lines.length;
    const ncol = lines[0].length;
    const grid: string[][] = [];

    for (let i = 0; i < nrow; i++) {
      grid[i] = Array(ncol).fill("");
      for (let j = 0; j < ncol; j++) {
        grid[i][j] = lines[i][j];
      }
    }

    this.grid = new Grid(grid);
  }

  part1() {
    for (let i = 0; i < this.grid.nrow; i++) {
      for (let j = 0; j < this.grid.ncol; j++) {
        const pt = new Point(i, j);
        const char = this.grid.getPVal(pt);
        if (this.seen.has(pt.id)) {
          continue;
        }
        this.dfs(pt, char, pt);
      }
    }
  }

  dfs(curP: Point, char: string, head: Point) {
    if (
      !this.grid.isPInside(curP) || this.grid.getPVal(curP) !== char ||
      this.seen.has(curP.id)
    ) {
      return;
    }
    this.seen.add(curP.id);
    if (!this.land.has(head.id)) {
      this.land.set(head.id, []);
    }
    this.land.get(head.id)!.push(curP);
    for (const dir of Dirs4) {
      const nextp = curP.move(dir);
      this.dfs(nextp, char, head);
    }
  }
  part2() {}
  res(): number {
    for (const value of this.land.values()) {
      const area = value.length;
      // const p = this.getP(key);
      // const char = this.grid.getPVal(p);
      const char = this.grid.getPVal(value[0]);
      const prem = this.getPrem(value, char);
      this.ans += area * prem;
    }
    return this.ans;
  }
  getPrem(list: Point[], char: string): number {
    let pre = 0;
    for (const p of list) {
      for (const dir of Dirs4) {
        const np = p.move(dir);
        if (!this.grid.isPInside(np) || this.grid.getPVal(np) !== char) {
          pre++;
        }
      }
    }
    return pre;
  }

  getP(s: string): Point {
    const p = s.split(":");
    return new Point(Number(p[0]), Number(p[1]));
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
