import { dirname, join } from "@std/path";

export class solution {
  strList: string[];
  nums: number[];
  memo = new Map<string, number>();
  ans = 0;

  constructor(input: string) {
    this.strList = input.split(" ");
    const nums = [];
    for (const str of this.strList) {
      nums.push(Number(str));
    }
    this.nums = nums;
  }
  part1() {
    for (const n of this.nums) {
      this.ans += this.dfs(n, 0, 75);
    }
  }

  dfs(num: number, cur: number, target: number): number {
    if (cur === target) {
      return 1;
    }

    const id = String(num) + ":" + String(cur);
    if (this.memo.has(id)) {
      return this.memo.get(id)!;
    }

    const nums = this.blink(num);
    let ans = 0;
    for (const n of nums) {
      ans += this.dfs(n, cur + 1, target);
    }
    this.memo.set(id, ans);
    return ans;
  }

  blink(stone: number): number[] {
    const newStones: number[] = [];
    if (stone === 0) {
      newStones.push(1);
    } else if (String(stone).length % 2 === 0) {
      const numStr = String(stone);
      const half = Math.floor(numStr.length / 2);
      const leftStr = numStr.slice(0, half);
      const rightStr = numStr.slice(half);

      // Convert parts directly to numbers
      const left = Number(leftStr);
      const right = Number(rightStr);

      newStones.push(left, right);
    } else {
      newStones.push(stone * 2024);
    }
    return newStones;
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
