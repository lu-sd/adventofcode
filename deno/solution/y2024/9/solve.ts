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
      let count = Number(this.input[i]);
      if (i % 2 === 0) {
        while (count--) {
          arr.push(id);
        }
        // arr.push(...Array(count).fill(id));
        id++;
      } else {
        while (count--) {
          arr.push(-1);
        }
        // arr.push(...Array(count).fill(-1));
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
      // [this.narr[l], this.narr[r]] = [this.narr[r], this.narr[l]];
      this.swap(l, r);
      l++;
      r--;
    }
  }

  findFile(id: number): [number, number] {
    let fLeft = -1;
    let len = 0;
    for (const [i, v] of this.narr.entries()) {
      if (v === id) {
        if (fLeft == -1) {
          fLeft = i;
        }
        len++;
      }
    }
    return [fLeft, len];
  }

  findFirstSpot(end: number, len: number): [number, boolean] {
    let curlen = 0;
    let start = -1;
    for (let i = 0; i < end; i++) {
      if (this.narr[i] === -1) {
        if (curlen === 0) {
          start = i;
        }
        curlen++;
        if (curlen === len) {
          return [start, true];
        }
        continue;
      }

      curlen = 0;
      start = -1;
    }
    return [-1, false];
  }

  part2() {
    for (let id = this.narr[this.narr.length - 1]; id >= 0; id--) {
      const [fileS, fLen] = this.findFile(id);
      const [freeS, find] = this.findFirstSpot(fileS, fLen);
      if (find) {
        this.swap2(freeS, fileS, fLen, id);
        // for (let i = 0; i < fLen; i++) {
        //   this.narr[fSart + i] = id;
        //   this.narr[fEdge + i] = -1;
        // }
      }
    }
  }
  swap2(freeS: number, fileS: number, len: number, fileid: number) {
    for (let idx = freeS; idx < freeS + len; idx++) {
      this.narr[idx] = fileid;
    }
    for (let idx = fileS; idx < fileS + len; idx++) {
      this.narr[idx] = -1;
    }
  }

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
