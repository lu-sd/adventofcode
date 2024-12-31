import { dirname, join } from "@std/path";

export class solve {
  input: string;
  strList: string[];
  ans = 0;
  constructor(input: string) {
    this.input = input;
    this.strList = this.input.split("\n");
  }

  // mul(23,4)
  // do() don't()
  mul2(str: string): number {
    let res = 0;
    let enable = true;
    for (let i = 0; i < str.length - 7; i++) {
      if (str.slice(i, i + 7) == "don't()") {
        enable = false;
      }
      if (str.slice(i, i + 4) == "do()") {
        enable = true;
      }
      if (str.slice(i, i + 4) == "mul(") {
        i += 4;
        const n1 = this.getNum(i, str);
        i += n1.numLength; // move 1 str at least
        if (str[i] != ",") {
          continue;
        }
        i++;
        const n2 = this.getNum(i, str);
        i += n2.numLength;
        if (str[i] != ")") {
          continue;
        }
        if (enable) {
          res += n1.value * n2.value;
        }
      }
    }
    return res;
  }
  mul(str: string): number {
    let res = 0;
    for (let i = 0; i < str.length - 4; i++) {
      if (str.slice(i, i + 4) != "mul(") {
        continue;
      }
      i += 4;
      const n1 = this.getNum(i, str);
      i += n1.numLength; // move 1 str at least
      if (str[i] != ",") {
        continue;
      }
      i++;
      const n2 = this.getNum(i, str);
      i += n2.numLength;
      if (str[i] != ")") {
        continue;
      }
      res += n1.value * n2.value;
    }
    return res;
  }

  getNum(idx: number, str: string): { value: number; numLength: number } {
    let value = 0;
    let numLength = 0;

    for (let i = idx; i < str.length - 1; i++) {
      const char = str[i];
      // const skip = str[i + 1];

      // Check if the character is a digit
      if (char >= "0" && char <= "9") {
        value = value * 10 + (char.charCodeAt(0) - "0".charCodeAt(0));
        numLength++;
        // if (skip == "," || skip == ")") {
        //   break;
        // }
      } else {
        return { value, numLength };
        // Return 0 and the index of the first non-numeric character
        // return [0, i - idx + 1];
      }
    }

    return { value, numLength };
    // If all characters are numeric, return the number and its length
    // return [num, numLength];
  }

  part1() {
    for (const str of this.strList) {
      this.ans += this.mul(str);
    }
  }
  part2() {
    this.ans = this.mul2(this.input);
  }
  res() {
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solve(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}
