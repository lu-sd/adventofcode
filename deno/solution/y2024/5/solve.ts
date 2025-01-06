import { dirname, join } from "@std/path";
import { IntsFromString } from "../../../lib/number.ts";

export class solution {
  ans = 0;
  adjList: Map<number, number[]>;
  manuals: number[][];

  constructor(input: string) {
    const lines = input.split("\n");
    const adjList = new Map<number, number[]>();
    const manuals: number[][] = [];
    let first = true;
    for (const line of lines) {
      const lineNum = IntsFromString(line);
      if (lineNum.length === 0) {
        first = false;
        continue;
      }
      if (first) {
        if (!adjList.has(lineNum[0])) {
          adjList.set(lineNum[0], []);
        }
        adjList.get(lineNum[0])!.push(lineNum[1]);
      } else {
        manuals.push(lineNum);
      }
    }
    this.adjList = adjList;
    this.manuals = manuals;
  }

  part1() {
    for (const manual of this.manuals) {
      if (this.isvalid(manual)) {
        this.ans += this.middle(manual);
      }
    }
  }
  isvalid(m: number[]): boolean {
    const seen = new Set();
    for (const item of m) {
      seen.add(item);
      if (!this.adjList.has(item)) {
        continue;
      }
      for (const child of this.adjList.get(item)!) {
        if (seen.has(child)) {
          return false;
        }
      }
    }
    return true;
  }

  middle(m: number[]) {
    return m[Math.floor(m.length / 2)];
  }

  part2() {
    for (const m of this.manuals) {
      if (!this.isvalid(m)) {
        this.ans += this.resorted(m);
      }
    }
  }

  resorted(m: number[]): number {
    const adj = this.buildAdj(m);
    const indegree = this.findDegree(m, adj);
    let ans = 0;
    ans += this.goodList(adj, indegree);
    return ans;
  }

  buildAdj(m: number[]): Map<number, number[]> {
    const adj = new Map<number, number[]>();
    for (const par of m) {
      const dependency = [];
      if (!this.adjList.has(par)) {
        continue;
      }
      for (const child of this.adjList.get(par)!) {
        if (m.includes(child)) {
          dependency.push(child);
        }
      }
      if (!adj.has(par)) {
        adj.set(par, []);
      }
      adj.get(par)!.push(...dependency);
    }
    return adj;
  }

  findDegree(m: number[], adj: Map<number, number[]>): Map<number, number> {
    const indegree = new Map<number, number>();
    for (const item of m) {
      indegree.set(item, 0);
    }
    for (const child of adj.values()) {
      for (const num of child) {
        indegree.set(num, indegree.get(num)! + 1);
      }
    }
    return indegree;
  }

  goodList(
    adj: Map<number, number[]>,
    indegree: Map<number, number>,
  ): number {
    const goodList: number[] = [];
    const queue: number[] = [];
    for (const [key, value] of indegree) {
      if (value === 0) {
        queue.push(key);
      }
    }

    while (queue.length > 0) {
      const head = queue.shift()!;
      goodList.push(head);
      for (const child of adj.get(head)!) {
        indegree.set(child, indegree.get(child)! - 1);
        if (indegree.get(child) == 0) {
          queue.push(child);
        }
      }
    }
    return goodList[Math.floor(goodList.length / 2)];
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
