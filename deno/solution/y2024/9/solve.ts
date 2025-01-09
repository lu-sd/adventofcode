import { dirname, join } from "@std/path";

export class solution {
  input: string;
  narr: number[];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    const arr: number[] = [];
    let id = 0;
    for (let i = 0; i < this.input.length; i++) {
      if (i % 2 === 0) {
        arr.push(...Array(Number(this.input[i])).fill(id));
        id++;
      } else {
        arr.push(...Array(Number(this.input[i])).fill(-1));
      }
    }
    this.narr = arr;
  }

  swap(l: number, r: number) {
    const temp = this.narr[l];
    this.narr[l] = this.narr[r];
    this.narr[r] = temp;
  }
  part1() {
    let l = 0;
    let r = this.narr.length - 1;
    while (l < r) {
      if (this.narr[l] !== -1) {
        l++;
        continue;
      }
      if (this.narr[r] === -1) {
        r--;
        continue;
      }
      [this.narr[l], this.narr[r]] = [this.narr[r], this.narr[l]];
      // this.swap(l, r);
      l++;
      r--;
    }
  }
  part2() {}
  res(): number {
    for (let i = 0; i < this.narr.length; i++) {
      const ele = this.narr[i];
      if (ele === -1) {
        continue;
      }
      this.ans += i * ele;
    }
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
