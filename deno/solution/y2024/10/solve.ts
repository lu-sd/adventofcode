import { dirname, join } from "@std/path";
import { Dirs4, Grid, Point } from "../../../lib/dataStructure.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;
  grid: Grid<number>;
  // seen: Map<string, boolean>;
  ansSet = new Set<string>();
  seen = new Set<string>();

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    const arr: number[][] = [];
    for (const [i, line] of this.lines.entries()) {
      arr[i] = Array(line.length);
      for (let j = 0; j < line.length; j++) {
        // arr[i][j] = Number(line[j]);
        arr[i][j] = line[j].charCodeAt(0) - "0".charCodeAt(0);
      }
    }
    this.grid = new Grid(arr);
  }
  part1() {
    for (let i = 0; i < this.grid.nrow; i++) {
      for (let j = 0; j < this.grid.ncol; j++) {
        if (this.grid.getVal(i, j) === 0) {
          const p = new Point(i, j);
          this.dfs1(p, p, 0);
        }
      }
    }
  }

  dfs1(start: Point, curP: Point, target: number) {
    if (
      !this.grid.isPInside(curP) || this.seen.has(curP.id) ||
      this.grid.getPVal(curP) !== target
    ) {
      return;
    }

    if (this.grid.getPVal(curP) === 9) {
      this.find(start.id, curP.id);
      return;
    }
    this.seen.add(curP.id);
    for (const dir of Dirs4) {
      this.dfs1(start, curP.move(dir), target + 1);
    }
    this.seen.delete(curP.id);
  }
  find(a: string, b: string) {
    this.ansSet.add(a + b);
  }

  part2() {
    for (let i = 0; i < this.grid.nrow; i++) {
      for (let j = 0; j < this.grid.ncol; j++) {
        if (this.grid.getVal(i, j) === 0) {
          const p = new Point(i, j);
          this.dfs2(p, 0);
        }
      }
    }
  }

  dfs2(curP: Point, target: number) {
    if (
      !this.grid.isPInside(curP) || this.seen.has(curP.id) ||
      this.grid.getPVal(curP) !== target
    ) {
      return;
    }

    if (this.grid.getPVal(curP) === 9) {
      this.ans++;
      return;
    }
    this.seen.add(curP.id);
    for (const dir of Dirs4) {
      this.dfs2(curP.move(dir), target + 1);
    }
    this.seen.delete(curP.id);
  }
  res1(): number {
    return this.ansSet.size;
  }

  res2(): number {
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solution(input);
  s1.part1();
  console.log("Part1 result ->", s1.res1());
  const s2 = new solution(input);
  s2.part2();
  console.log("Part2 result ->", s2.res2());
}
