import { dirname, join } from "@std/path";
import { Grid } from "../../../lib/dataStructure.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;
  grid: Grid<string>;
  ans2 = 0;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    // const array: string[][] = [];
    // for (const [r, line] of this.lines.entries()) {
    //   array[r] = Array(line.length);
    //   for (let c = 0; c < line.length; c++) {
    //     const char = line[c];
    //     array[r][c] = char;
    //   }
    // }
    const array: string[][] = this.lines.map((line) => [...line]);
    this.grid = new Grid(array);
  }
  part1() {
    const target = "XMAS";
    const dirs8 = [];
    for (let x = -1; x <= 1; x++) {
      for (let y = -1; y <= 1; y++) {
        if (x == 0 && y == 0) {
          continue;
        }
        dirs8.push([x, y]);
      }
    }
    for (let r = 0; r < this.grid.nrow; r++) {
      for (let c = 0; c < this.grid.ncol; c++) {
        for (const dir of dirs8) {
          let valid = true;
          for (let i = 0; i < target.length; i++) {
            const nr = r + i * dir[0];
            const nc = c + i * dir[1];
            if (
              !this.grid.isInside(nr, nc) ||
              this.grid.getVal(nr, nc) !== target[i]
            ) {
              valid = false;
              break;
            }
          }
          if (valid) {
            this.ans++;
          }
        }
      }
    }
  }
  part2() {
    const dirs = [
      [[1, 1], [-1, -1]],
      [[-1, 1], [1, -1]],
    ];

    const points = new Map([
      ["M", 1],
      ["S", 2],
    ]);

    for (let r = 0; r < this.grid.nrow; r++) {
      for (let c = 0; c < this.grid.ncol; c++) {
        const char = this.grid.getVal(r, c);
        let valid = true;
        if (char !== "A") {
          continue;
        }
        for (const dir of dirs) {
          let counter = 0;
          for (const side of dir) {
            const nr = r + side[0];
            const nc = c + side[1];

            if (!this.grid.isInside(nr, nc)) {
              valid = false;
              break;
            }
            const nChar = this.grid.getVal(nr, nc);
            if (points.has(nChar)) {
              counter += points.get(nChar)!;
            }
          }
          if (counter !== 3) {
            valid = false;
            break;
          }
        }
        if (valid) {
          this.ans2++;
        }
      }
    }
  }
  res(): number {
    return this.ans;
  }
  res2(): number {
    return this.ans2;
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
