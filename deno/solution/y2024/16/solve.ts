import { dirname, join } from "@std/path";
import { Grid, Point } from "../../../lib/dataStructure.ts";
import { Dirs4 } from "../../../lib/dataStructure.ts";
type Node = {
  pos: Point;
  dir: Point;
  score: number;
};
export class solution {
  input: string;
  lines: string[];
  grid: Grid<string>;
  minScore = new Map<string, number>();
  ans = 0;
  start: Point;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    // const array: string[][] = [];
    // for (const [i, line] of this.lines.entries()) {
    //   array[i] = Array(line.length);
    //   for (let j = 0; j < line.length; j++) {
    //     array[i][j] = line[j];
    //   }
    // }
    // also you can use map and split to let loop more consise
    const array: string[][] = this.lines.map((line) => line.split(""));
    this.grid = new Grid(array);
    this.start = new Point(this.grid.nrow - 2, 1);
  }
  part1() {
    this.bfs(this.start, Dirs4[1]);
  }

  bfs(p: Point, dir: Point) {
    const queue: Node[] = [];
    queue.push({ pos: p, dir: dir, score: 0 });
    this.minScore.set(p.id, 0);
    while (queue.length > 0) {
      const n = queue.length;
      for (let i = 0; i < n; i++) {
        const top = queue.shift()!;
        if (this.grid.getPVal(top.pos) === "E") {
          if (this.ans === 0 || this.ans > top.score) {
            this.ans = top.score;
          }
        }

        for (const d of Dirs4) {
          const nextP = top.pos.move(d);
          if (this.grid.getPVal(nextP) === "#") {
            continue;
          }
          let turn = 0;
          if (!top.dir.equals(d)) {
            turn = 1;
          }
          const nextS = top.score + 1 + turn * 1000;
          if (
            !this.minScore.has(nextP.id) || nextS < this.minScore.get(nextP.id)!
          ) {
            this.minScore.set(nextP.id, nextS);
            queue.push({ pos: nextP, dir: d, score: nextS });
          }
        }
      }
    }
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
