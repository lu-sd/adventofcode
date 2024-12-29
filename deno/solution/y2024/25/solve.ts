import { dirname, join } from "@std/path";
type pairs = Map<string, number[][]>;
export class solution {
  input: string;
  ans = 0;

  constructor(input: string) {
    this.input = input;
  }

  findMatch(heights: pairs) {
    for (const pair1 of heights.get("#")!) {
      for (const pair2 of heights.get(".")!) {
        if (this.isvalid(pair1, pair2)) {
          this.ans++;
        }
      }
    }
  }

  isvalid(p1: number[], p2: number[]) {
    for (const [idx, v] of p1.entries()) {
      if (v + p2[idx] > 5) {
        return false;
      }
    }
    return true;
  }

  getHeights(grid: string[]): number[] {
    const ncol = grid[0].length;
    const nrow = grid.length;
    const res = new Array(ncol).fill(0);
    for (let c = 0; c < ncol; c++) {
      let count = -1;

      for (let r = 0; r < nrow; r++) {
        const char = grid[r][c];
        if (char === "#") {
          count++;
        }
      }

      res[c] = count;
    }
    return res;
  }

  buildGrid(lines: string[]): pairs {
    const res: pairs = new Map();
    let hint = "0";

    const grid: string[] = [];
    for (const line of lines) {
      // if (hint == "0" && line[0] == "#") {
      //   hint = "#";
      // }
      // if (hint == "0" && line[0] == ".") {
      //   hint = ".";
      // }
      if (hint == "0") {
        hint = line[0];
      }

      if (line.length == 0) {
        const height = this.getHeights(grid);
        if (!res.has(hint)) {
          res.set(hint, []);
        }
        res.get(hint)!.push(height);
        // reset everything
        hint = "0";
        grid.length = 0;
        continue;
      }

      grid.push(line);
    }

    const height = this.getHeights(grid);
    res.get(hint)!.push(height);
    return res;
  }

  part1() {
    const lines = this.input.split("\n");
    const heights = this.buildGrid(lines);
    this.findMatch(heights);
  }
  part2() {
    this.ans = 0;
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
