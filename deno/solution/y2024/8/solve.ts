import { dirname, join } from "@std/path";
import { Point } from "../../../lib/dataStructure.ts";

export class solution {
  input: string;
  lines: string[];
  anti: Map<string, Point[]>;
  antiNode = new Set<string>();
  dirs = [1, -1];
  nRow: number;
  nCol: number;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    const nRow = this.lines.length;
    const nCol = this.lines[0].length;
    const anti = new Map<string, Point[]>();
    for (let i = 0; i < nRow; i++) {
      for (let j = 0; j < nCol; j++) {
        const char = this.lines[i][j];
        if (char !== ".") {
          //const points = anti.get(char) ?? [];
          // points.push(new Point(i, j));
          // anti.set(char, points);
          if (!anti.has(char)) {
            anti.set(char, []);
          }
          anti.get(char)!.push(new Point(i, j));
        }
      }
    }
    this.anti = anti;
    this.nRow = nRow;
    this.nCol = nCol;
  }

  part1() {
    for (const points of this.anti.values()) {
      for (let i = 0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
          const ptParis = [points[i], points[j]];
          this.findNode(ptParis);
        }
      }
    }
  }

  findNode(ptParis: Point[]) {
    const [dx, dy] = ptParis[0].dist(ptParis[1]);
    for (const [i, pt] of ptParis.entries()) {
      const npt = new Point(pt.x + dx * this.dirs[i], pt.y + dy * this.dirs[i]);
      if (this.isInside(npt)) {
        this.antiNode.add(npt.id);
      }
    }
  }

  findNode2(ptParis: Point[]) {
    const [dx, dy] = ptParis[0].dist(ptParis[1]);
    for (const [i, pt] of ptParis.entries()) {
      this.antiNode.add(pt.id);
      let step = 1;
      while (true) {
        const npt = new Point(
          pt.x + dx * step * this.dirs[i],
          pt.y + dy * step * this.dirs[i],
        );

        if (this.isInside(npt)) {
          this.antiNode.add(npt.id);
          step++;
        } else {
          break;
        }
      }
    }
  }
  isInside(pt: Point): boolean {
    return pt.x < this.nRow && pt.x >= 0 && pt.y < this.nCol && pt.y >= 0;
  }

  part2() {
    for (const points of this.anti.values()) {
      for (let i = 0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
          const ptParis = [points[i], points[j]];
          this.findNode2(ptParis);
        }
      }
    }
  }
  res(): number {
    return this.antiNode.size;
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
