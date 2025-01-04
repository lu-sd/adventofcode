import { dirname, join } from "@std/path";

export class solution {
  lines: string[];
  nrow: number;
  ncol: number;
  grid: string[][];
  dirs: number[][];
  seen = new Set<string>();
  start = { r: 0, c: 0 };

  constructor(input: string) {
    this.lines = input.split("\n");
    this.nrow = this.lines.length;
    this.ncol = this.lines[0].length;
    this.start = { r: 0, c: 0 };
    const grid: string[][] = [];
    for (const [r, line] of this.lines.entries()) {
      grid.push(new Array(line.length).fill(""));
      for (let c = 0; c < line.length; c++) {
        const char = line[c];
        grid[r][c] = char;
        if (char === "^") {
          this.start = { r: r, c: c };
        }
      }
    }
    const dirs = [
      [-1, 0],
      [0, 1],
      [1, 0],
      [0, -1],
    ];

    this.grid = grid;
    this.dirs = dirs;
  }
  isInside(r: number, c: number) {
    return r < this.nrow && r >= 0 && c < this.ncol && c >= 0;
  }
  getId(r: number, c: number) {
    // return String(r) + ":" + String(c);
    return `${r}:${c}`;
  }
  getId2(pt: { r: number; c: number }, dir: number[]): string {
    return `${pt.r}:${pt.c}:${dir}`;
  }
  test(r: number, c: number) {
    return String(r) + ":" + String(c);
    // return `${r}:${c}`;
  }
  part1() {
    for (let r = 0; r < this.nrow; r++) {
      for (let c = 0; c < this.ncol; c++) {
        if (this.grid[r][c] === "^") {
          this.dfs(r, c, 0);
        }
      }
    }
  }
  dfs(r: number, c: number, dir: number) {
    const id = this.getId(r, c);
    this.seen.add(id);
    const curdir = this.dirs[dir % 4];
    const nr = r + curdir[0];
    const nc = c + curdir[1];
    if (!this.isInside(nr, nc)) {
      return;
    }
    if (this.grid[nr][nc] !== "#") {
      this.dfs(nr, nc, dir);
    } else {
      this.dfs(r, c, dir + 1);
    }
  }
  part2() {
    let ans = 0;
    for (let r = 0; r < this.nrow; r++) {
      for (let c = 0; c < this.ncol; c++) {
        if (this.grid[r][c] === ".") {
          const visited = new Map<string, boolean>();
          this.grid[r][c] = "#";
          ans += this.dfs2(this.start, 0, visited);
          this.grid[r][c] = ".";
        }
      }
    }
    return ans;
  }
  dfs2(
    pt: { r: number; c: number },
    dir: number,
    visited: Map<string, boolean>,
  ): number {
    const curdir = this.dirs[dir % 4];
    const nextPt = { r: pt.r + curdir[0], c: pt.c + curdir[1] };
    if (!this.isInside(nextPt.r, nextPt.c)) {
      return 0;
    }
    if (this.grid[nextPt.r][nextPt.c] === "#") {
      if (visited.get(this.getId2(nextPt, curdir))) {
        return 1;
      }
      visited.set(this.getId2(nextPt, curdir), true);
      return this.dfs2(pt, dir + 1, visited);
    } else {
      return this.dfs2(nextPt, dir, visited);
    }
  }
  res(): number {
    return this.seen.size;
  }

  res2(): number {
    return this.part2();
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
  console.log("Part2 result ->", s2.res2());
}
